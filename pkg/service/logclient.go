package service

import (
	"bytes"
	"encoding/json"
	"github.com/1uvu/bitlog/pkg/common"
	"github.com/1uvu/bitlog/pkg/config"
	"github.com/1uvu/bitlog/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

type (
	LogClient struct {
		sync.Mutex // 用来控制 seq 计数, 只有 reportLog 和 gotWriteEvent 日志入队时存在并发问题

		*fsnotify.Watcher

		conf       *config.CollectorConfig
		now        time.Time // 用于获取当前 log 的 day
		seqNum     int64     // 记录当前 log 的序号
		serverAddr string    // 日志接受服务地址

		stopCh <-chan struct{}
	}
)

const (
	RETRY_COUNT = 3
)

// NewLogClient TODO 支持动态获取 serverAddr
//
//	当前直接从配置中读取
func NewLogClient(conf *config.CollectorConfig, stopCh <-chan struct{}) (*LogClient, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	logClient := &LogClient{
		Watcher:    watcher,
		now:        now,
		seqNum:     1,
		serverAddr: conf.LogServer.Address,
		stopCh:     stopCh,
	}

	err = logClient.Add(utils.CurrentDayLogFilepath(conf.Base.BasePath, conf.Node.LoggerName, now))
	if err != nil {
		return nil, err
	}
	return logClient, nil
}

func (c *LogClient) Run() {
	go c.watchLog()
	c.runAndWait()
}

func (c *LogClient) runAndWait() {
	<-c.stopCh
	// 清理 watcher
	for _, w := range c.WatchList() {
		c.Remove(w)
	}
	c.Close()
}

func (c *LogClient) watchLog() {
	for {
		select {
		case event, ok := <-c.Events:
			if ok {
				// 判断是否是日志写事件
				if event.Op == fsnotify.Write {
					c.gotWriteEvent(event)
				}
			}
		case <-c.stopCh:
			return
		case err := <-c.Errors:
			if err != nil {
				log.Printf("[watchLog]%v", err)
			}
		}
	}
}

// gotWriteEvent TODO fix
// - 日志写入快, 读日志由于阻塞导致事件处理不及时, 使得读到的最后一行日志是重复的（全是最新的那一行）
// - 由于 http 调用延时导致事件处理不及时, 使得读到的最后一行日志是重复的（全是最新的那一行）: 开协程来处理
func (c *LogClient) gotWriteEvent(event fsnotify.Event) {
	// 读日志
	var data []byte
	data, err := utils.ReadLastLine(event.Name)
	if err != nil {
		log.Printf("[gotWriteEvent]read log err:%v", err)
	}

	// test
	log.Printf("[gotWriteEvent]event:%v, data:%s", event, string(data))

	if string(data) != common.LOG_EOF {
		go c.reportLog(data)
	} else {
		// 如果读到的是 eof 跳过入队, 并 refreshWatchFile
		err := c.refreshWatchFile()
		if err != nil {
			log.Printf("[gotWriteEvent]refresh err:%v", err)
		}
	}
}

// reportLog 调用 http 接口, 不返回错误, 调用服务仅做简单的重试
// 开协程来调用此方法, 此方法不会死循环, 不存在泄露风险
func (c *LogClient) reportLog(data []byte) {
	reportReq := LogReportRequest{
		Data:   data,
		SeqNum: c.seqNum,
	}
	reqAsBytes, err := json.Marshal(reportReq)
	if err != nil {
		log.Printf("[reportLog]json err:%v", err)
	}

	// Test
	log.Printf("report log req:%+v", reportReq)

	var resp *http.Response
	for i := 0; i < RETRY_COUNT; i++ {
		resp, err = http.Post(c.serverAddr, "", bytes.NewReader(reqAsBytes))
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("[reportLog]post resp:%v,err:%v", resp, err)
	}

	// 更新 seq
	c.Lock()
	c.seqNum = c.seqNum + 1
	c.Unlock()
}

// refreshWatchFile	当从当前 watch log 文件读取到 EOF 时,
// 调用这个函数来刷新 watch 的 log 文件
func (c *LogClient) refreshWatchFile() error {
	// 循环将 now ＋ 1 day, 直到日志文件存在, 然后 watch
	// 阈值是当前模块日志文件夹的文件数量
	var (
		i               int
		logfileCount    = utils.DirFileCount(filepath.Join(c.conf.Base.BasePath, c.conf.Node.LoggerName))
		logfilePath     string
		nextLogfilePath string
	)
	for nextLogfilePath = utils.CurrentDayLogFilepath(c.conf.Base.BasePath, c.conf.Node.LoggerName, c.now.Add(24*time.Hour)); i < logfileCount && !utils.IsFileExisted(nextLogfilePath); {
		c.now = c.now.Add(24 * time.Hour)
		i++
	}
	if len(c.WatchList()) > 0 {
		logfilePath = c.WatchList()[0]
		err := c.Remove(logfilePath)
		if err != nil {
			log.Printf("[refreshWatchFile]%v", err)
		}
	}
	return c.Add(nextLogfilePath)
}
