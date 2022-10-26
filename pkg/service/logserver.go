package service

import (
	"encoding/json"
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
		Data   []byte `json:"data"`
		SeqNum int64  `json:"seq_num"` // 当前传送的 logtree 的序号，防止乱序
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
	dataAsBytes, err := io.ReadAll(body)
	if err != nil {
		// TODO 先 panic
		panic(fmt.Sprintf("[reportLogHandleFunc]read body err:%v", err))
	}

	var reportReq LogReportRequest
	err = json.Unmarshal(dataAsBytes, &reportReq)
	if err != nil {
		// TODO 先 panic
		panic(fmt.Sprintf("[reportLogHandleFunc]json err:%v", err))
	}

	// TODO
	// 	- 1. 这里不解决乱序问题，由插件自行解决
	// 	- 2. 传给 loop 插件
	c.String(http.StatusOK, "OK")

	log.Println(reportReq.SeqNum)
	log.Printf("%s\n", string(reportReq.Data))
}
