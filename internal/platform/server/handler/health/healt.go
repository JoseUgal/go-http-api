package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  CheckHandler returns HTTP handler to perdorm health checks.
func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context ) {
		ctx.String(http.StatusOK, "Everything is OK!")
	}
}