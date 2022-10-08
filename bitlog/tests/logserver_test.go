package tests

import (
	"github.com/1uvu/bitlog/collector/logserver"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestLogServer(t *testing.T) {
	r := gin.Default()
	logserver.Register(r)

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	r.Run(":8080")
}
