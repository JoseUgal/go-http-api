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
	"github.com/kelseyhightower/envconfig"
)

func Run() error {

	var cfg config 
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	fmt.Println(cfg)

	mysqlURI := fmt.Sprintf("%s:%v@/%s", cfg.dbUser, cfg.dbPort, cfg.dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	// Create CommandBus
	var (
		commandBus = inmemory.NewCommandBus()
		eventBus = inmemory.NewEventBus()
	)

	
	courseRepository := mysql.NewCourseRepository(db, cfg.dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)


	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout , commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host			string			`default:"localhost"`
	Port			uint			`default:"8080"`
	ShutdownTimeout	time.Duration	`default:"10s"`
	// Database configuration
	dbUser			string			`default:"root"`
	dbPass			string			
	dbHost			string			`default:"localhost"`
	dbPort			uint			`default:"3306"`
	dbName			string			`default:"mooc"`
	dbTimeout		time.Duration	`default:"5s"`
}