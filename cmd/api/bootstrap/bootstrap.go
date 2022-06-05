package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/JoseUgal/go-http-api/internal/creating"
	"github.com/JoseUgal/go-http-api/internal/platform/bus/inmemory"
	"github.com/JoseUgal/go-http-api/internal/platform/server"
	"github.com/JoseUgal/go-http-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

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

	srv := server.New(host, port, creatingCourseService)
	return srv.Run()
}