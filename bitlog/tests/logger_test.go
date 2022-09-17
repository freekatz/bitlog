package tests

import (
	"testing"

	"github.com/1uvu/bitlog/pkg/common"
)

func TestGetCMDLogger(t *testing.T) {
	cmdLogger := common.GetLogger("test_cmd", "", false)
	cmdLogger.Info("test cmd logger")
}

func TestGetFileLogger(t *testing.T) {
	fileLogger := common.GetLogger("test_file", "./tmp", true)
	fileLogger.Info("test file logger")
}
