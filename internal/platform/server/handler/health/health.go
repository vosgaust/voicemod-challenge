package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Status string `json:"status"`
}

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response{"ok"})
	}
}
