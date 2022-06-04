package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/JoseUgal/go-http-api/internal/platform/server"
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

	mysqlURI := fmt.Sprintf("%s:%s@tcp/(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	_, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	// courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, nil)
	return srv.Run()
}