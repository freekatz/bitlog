package common

import (
	"os"

	"github.com/1uvu/bitlog/pkg/errorx"
)

type EnvKey int

const (
	ROOT_DIR = iota
	CLIENT_DIR
	LOG_DIR
	CONFIG_DIR
	COLLECTOR_CONFIG_NAME
)

var envKeyStr = map[EnvKey]string{
	ROOT_DIR:              "BITLOG_ROOT_DIR",
	CLIENT_DIR:            "BITLOG_CLIENT_DIR",
	LOG_DIR:               "BITLOG_LOG_DIR",
	CONFIG_DIR:            "BITLOG_CONFIG_DIR",
	COLLECTOR_CONFIG_NAME: "BITLOG_COLLECTOR_CONFIG_NAME",
}

func (envKey EnvKey) String() string {
	return envKeyStr[envKey]
}

// LookupEnvPairs 根据所有 key 查找环境变量，如果中途有 error 会直接中断
func LookupEnvPairs(envPairs *map[string]string) error {
	for keyStr := range *envPairs {
		env, ok := os.LookupEnv(keyStr)
		if !ok {
			return errorx.ErrEnvLookupFailed
		}
		(*envPairs)[keyStr] = env
	}
	return nil
}

// LookupEnvPairsByKey 根据所有 key 查找环境变量，如果中途有 error 会直接中断
// 通过限定 EnvKey 使得调用者只能从已定义的 Env 中选取
func LookupEnvPairsByKey(envPairs *map[EnvKey]string) error {
	for key := range *envPairs {
		keyStr := key.String()
		if keyStr == "" {
			return errorx.ErrEnvKeyInvalid
		}
		env, ok := os.LookupEnv(keyStr)
		if !ok {
			return errorx.ErrEnvLookupFailed
		}
		(*envPairs)[key] = env
	}
	return nil
}
