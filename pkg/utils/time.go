package utils

import (
	"path/filepath"
	"time"
)

const (
	TIME_LAYOUT_DEFAULT = "2006-01-02 15:04:05.000"
	TIME_LAYOUT_SECOND  = "2006-01-02-15-04-05"
	TIME_LAYOUT_DAY     = "2006-01-02"
	TIME_LOCATION       = "Asia/Shanghai"
)

func TimeStr(t time.Time) string {
	return t.Format(TIME_LAYOUT_DEFAULT)
}

func TimeStrLocal(t time.Time) string {
	timeStr := TimeStr(t)
	location, _ := time.LoadLocation(TIME_LOCATION)
	localTime, _ := time.ParseInLocation(TIME_LAYOUT_DEFAULT, timeStr, location)
	return localTime.String()
}

func CurrentDay(now time.Time) string {
	return now.Format(TIME_LAYOUT_DAY)
}

func CurrentDayLogFilename(now time.Time) string {
	return CurrentDay(now) + ".log"
}

// CurrentDayLogFilepath basePath/loggerName/day.log
func CurrentDayLogFilepath(basePath, loggerName string, now time.Time) string {
	return filepath.Join(basePath, loggerName, CurrentDayLogFilename(now))
}
