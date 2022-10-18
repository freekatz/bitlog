package tests

import (
	"github.com/1uvu/bitlog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestLogServer(t *testing.T) {
	server := service.NewLogServer(gin.Default())

	server.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	server.Run(":8080")
}
