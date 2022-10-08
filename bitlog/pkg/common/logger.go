package common

import (
	"fmt"
	"github.com/1uvu/bitlog/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"sync"
)

/*
logger
	实现一个日志库，实现自动按日志输出的 day 来划分日志文件
	日志输出格式就是最简单的：timestamp + level + name + data
*/

type (
	Logger interface {
		Info(string, ...interface{})
		Warn(string, ...interface{})
		Error(string, ...interface{})
		Fatal(string, ...interface{})
	}

	defaultLogger struct {
		basePath string

		loggerName string
		day        string
		filepath   string // filepath = basePath + "/" + loggerName + day + ".log"

		log  log.Logger
		logf *os.File
	}
)

const (
	LOG_EOF = "[EOF]"
)

var (
	loggers map[string]Logger
	mux     sync.RWMutex
)

func init() {
	loggers = map[string]Logger{}
}

func GetLogger(loggerName, basePath string) Logger {
	mux.Lock()
	if _, ok := loggers[loggerName]; !ok {
		loggers[loggerName] = newLogger(loggerName, basePath)
	}
	mux.Unlock()

	mux.RLock()
	defer mux.RUnlock()
	return loggers[loggerName]
}

func newLogger(loggerName, basePath string) Logger {
	// new logger
	l := &defaultLogger{
		basePath:   basePath,
		loggerName: loggerName,
		log:        log.Logger{},
	}
	l.initLogger()

	return l
}

func (l *defaultLogger) Info(format string, msg ...interface{}) {
	l.info(fmt.Sprintf(format, msg...))
}
func (l *defaultLogger) info(msg string) {
	l.printf("[I][%s]:%s", l.loggerName, msg)
}

func (l *defaultLogger) Warn(format string, msg ...interface{}) {
	l.warn(fmt.Sprintf(format, msg...))
}
func (l *defaultLogger) warn(msg string) {
	l.printf("[W][%s]:%s", l.loggerName, msg)
}

func (l *defaultLogger) Error(format string, msg ...interface{}) {
	l.error(fmt.Sprintf(format, msg...))
}

func (l *defaultLogger) error(msg string) {
	l.printf("[E][%s]:%s", l.loggerName, msg)
}

func (l *defaultLogger) Fatal(format string, msg ...interface{}) {
	l.fatal(fmt.Sprintf(format, msg...))
}
func (l *defaultLogger) fatal(msg string) {
	l.printf("[F][%s]:%s", l.loggerName, msg)
}

func (l *defaultLogger) printf(format string, msg ...interface{}) {
	// 检查是否需要切换新的日志文件
	l.initLogger()
	l.log.Printf(format, msg...)
}

func (l *defaultLogger) initLogger() {
	// 如果当前 day 与 logger day 一样，无需初始化新的 log 文件
	if l.day == utils.CurrentDay() {
		return
	}
	// 关闭旧的 log 文件，关闭前会写入一个 EOF 日志用来标记当前日志文件的结束
	if l.logf != nil {
		// EOF
		l.log.Printf(LOG_EOF)
		l.logf.Close()
	}

	l.day = utils.CurrentDay()
	l.filepath = filepath.Join(l.basePath, l.loggerName+"-"+l.day+".log")

	// set log writer
	f, err := os.OpenFile(l.filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("[initLogger]%v", err))
	}
	l.logf = f
	l.log.SetOutput(l.logf)
}
