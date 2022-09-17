package utils

import (
	"time"
)

const (
	TIME_LAYOUT    = "2006-01-02 15:04:05.000"
	TIME_LOCALTION = "Asia/Shanghai"
)

func TimeStr(t time.Time) string {
	return t.Format(TIME_LAYOUT)
}

func TimeStrLocal(t time.Time) string {
	timeStr := TimeStr(t)
	localtion, _ := time.LoadLocation(TIME_LOCALTION)
	localTime, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, localtion)
	return localTime.String()
}
