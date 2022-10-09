package logserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type (
	LogServer struct {
		*gin.Engine
	}
	LogReportRequest struct {
		Data  [][]byte `json:"data"` // 打包传输的 log 数组
		Start int64    `json:"start"`
		End   int64    `json:"end"`
		// TODO 支持处理压缩的 log
	}
)

func NewLogServer(r *gin.Engine) *LogServer {
	server := &LogServer{
		Engine: r,
	}
	server.POST("/", reportLogHandleFunc)
	return server
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
