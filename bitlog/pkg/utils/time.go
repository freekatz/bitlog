package utils

import (
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

func CurrentDay() string {
	return time.Now().Format(TIME_LAYOUT_DAY)
}
