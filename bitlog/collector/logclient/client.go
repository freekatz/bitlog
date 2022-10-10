package logclient

import (
	"bytes"
	"encoding/json"
	"github.com/1uvu/bitlog/collector/logserver"
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
		sync.Mutex // 用来控制 start, end 计数的并发, 只有 reportLog 和 gotWriteEvent 日志入队时存在并发问题

		*fsnotify.Watcher

		conf       *config.CollectorConfig
		now        time.Time // 用于获取当前 log 的 day
		logQueue   *LogQueue // 线程安全的先入先出队列, 限定容量 24, 到达容量限制后会打包上报并重置
		start, end int64     // 记录当前 queue 中 log 的范围序号
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
		logQueue:   &LogQueue{},
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
	go c.timeWindow()
	go c.watchLog()
	c.runAndWait()
}

func (c *LogClient) runAndWait() {
	for {
		select {
		case <-c.stopCh:
			// 关闭前上报一次
			if c.start < c.end {
				c.reportLog()
			}
			// 清理 watcher
			for _, w := range c.WatchList() {
				c.Remove(w)
			}
			c.Close()
			return
		}
	}
}

func (c *LogClient) timeWindow() {
	// 周期计时器, 10s, 每次触发会将 logQueue 的数据打包上报
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			// 如果队列是空的, 则跳过
			if c.logQueue.Len() != 0 {
				// 如果不空则打包上报
				go c.reportLog()
			}
		case <-c.stopCh:
			ticker.Stop()
			return
		}
	}
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
		// 入队
		exceeded := c.logQueue.Enqueue(data)
		// 更新 end
		c.Lock()
		c.end += 1
		c.Unlock()
		// 如果入队返回已满, 则 reportLog
		if exceeded {
			go c.reportLog()
		}
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
func (c *LogClient) reportLog() {
	data := c.logQueue.Pack()
	reportReq := logserver.LogReportRequest{
		Data:  data,
		Start: c.start,
		End:   c.end,
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

	// 更新 start
	c.Lock()
	c.start = c.end + 1
	c.end = c.end + 1
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
