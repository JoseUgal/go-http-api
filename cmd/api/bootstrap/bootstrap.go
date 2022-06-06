package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/JoseUgal/go-http-api/internal"
	"github.com/JoseUgal/go-http-api/internal/creating"
	"github.com/JoseUgal/go-http-api/internal/increasing"
	"github.com/JoseUgal/go-http-api/internal/platform/bus/inmemory"
	"github.com/JoseUgal/go-http-api/internal/platform/server"
	"github.com/JoseUgal/go-http-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080
	shutdownTimeout = 10 * time.Second

	dbUser = "root"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "mooc"
	dbTimeout = 5 * time.Second
)

func Run() error {

	mysqlURI := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	// Create CommandBus
	var (
		commandBus = inmemory.NewCommandBus()
		eventBus = inmemory.NewEventBus()
	)

	
	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)


	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout , commandBus)
	return srv.Run(ctx)
}