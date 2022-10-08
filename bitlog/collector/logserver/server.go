package logserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type (
	ReportLogRequest struct {
		Data  [][]byte `json:"data"`
		Start int64    `json:"start"`
		End   int64    `json:"end"`
	}
)

func Register(r *gin.Engine) {
	r.POST("/", reportLogHandleFunc)
}

func reportLogHandleFunc(c *gin.Context) {
	body := c.Request.Body
	data, err := io.ReadAll(body)
	if err != nil {
		// TODO 先 panic
		panic(fmt.Sprintf("[reportLogHandleFunc]%v", err))
	}
	// TODO
	// 	- 1. 解决乱序问题
	// 	- 2. 传给 loop 插件
	c.String(http.StatusOK, "OK")
	log.Println(string(data))
}
