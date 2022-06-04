package courses

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	mooc "github.com/JoseUgal/go-http-api/internal"
	"github.com/gin-gonic/gin"
)

// The struct defines the properties that you need to perform the query.
// This properties includes in request Payload.
type createRequest struct {
	ID		 string `json:"id" binding:"required"`
	Name	 string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context){
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		// if err := Save(ctx, course); err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, err.Error())
		// }

		fmt.Println(course)
		fmt.Println(course.Name(), course.Duration())
		ctx.Status(http.StatusCreated)
	}
}

const (
	dbUser = "jugal"
	dbPass = "jugal"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "jugal"
)

// Save persist the course on the database.
func Save(ctx context.Context, course mooc.Course) error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp/(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if( err != nil ){
		return err
	}

	type sqlCourse struct {
		ID		 string `db:"id"`
		Name 	 string `db:"name"`
		Duration string `db:"duration"`
	}

	fmt.Println(db)

	return nil
}