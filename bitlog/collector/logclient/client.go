package logclient

import (
	"github.com/1uvu/bitlog/pkg/config"
	"github.com/1uvu/bitlog/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"log"
	"path/filepath"
	"time"
)

type (
	LogClient struct {
		*fsnotify.Watcher

		conf *config.CollectorConfig
		now  time.Time // 用于获取当前 log 的 day

		logQueue *LogQueue // 环形队列, 限定容量 24, 到达容量限制后会打包上报并重置

		serverAddr string // 日志接受服务地址
	}
)

// NewLogClient TODO 支持动态获取 serverAddr
//
//	当前直接从配置中读取
func NewLogClient(conf *config.CollectorConfig) (*LogClient, error) {
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
	}

	err = logClient.Add(utils.CurrentDayLogFilepath(conf.Base.BasePath, conf.Node.LoggerName, now))
	if err != nil {
		return nil, err
	}
	return logClient, nil
}

func (c *LogClient) Run() error {
	// TODO 运行 watcher
	go c.timeWindow()
	return nil
}

func (c *LogClient) timeWindow() {
	// 周期计时器, 10s, 每次触发会将 logQueue 的数据打包上报
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			// 如果队列是空的, 则跳过
			if c.logQueue.Len() == 0 {
				return
			}
			// 如果不空则打包上报
			c.ReportLog()
		}
	}
}

// ReportLog 不返回错误, 调用服务仅做简单的重试
func (c *LogClient) ReportLog() {
	//data := c.logQueue.Pack()

}

// RefreshWatch	当从当前 watch log 文件读取到 EOF 时,
//	调用这个函数来刷新 watch 的 log 文件
func (c *LogClient) RefreshWatch() error {
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
			// TODO log err and continue, not return
			log.Printf("[RefreshWatch]%v", err)
		}
	}
	return c.Add(nextLogfilePath)
}
