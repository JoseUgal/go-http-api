package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JoseUgal/go-http-api/internal/creating"
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
	)

	
	courseRepository := mysql.NewCourseRepository(db)
	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout , commandBus)
	return srv.Run(ctx)
}