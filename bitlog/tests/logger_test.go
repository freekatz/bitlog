package tests

import (
	"github.com/1uvu/bitlog/pkg/common"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	cmdLogger := common.GetLogger("test_logger", "")
	cmdLogger.Info("test info")
	time.Sleep(time.Second)
	cmdLogger.Warn("test warn")
	cmdLogger.Error("test error")
}
