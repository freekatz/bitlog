package config

import (
	"github.com/1uvu/bitlog/pkg/errorx"

	"github.com/spf13/viper"
)

type (
	CollectorConfig struct {
		Base      *BaseConfig      `mapstructure:"base"`
		Node      *NodeConfig      `mapstructure:"node"`
		LogClient *LogClientConfig `mapstructure:"logclient"`
		LogServer *LogServerConfig `mapstructure:"logserver"`
	}
	BaseConfig struct {
		BasePath string `mapstructure:"basepath"`
	}
	NodeConfig struct {
		RPC        *RPCConfig `mapstructure:"rpc"`
		LoggerName string     `mapstructure:"loggername"`
	}
	RPCConfig struct {
		Address  string `mapstructure:"address"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
	LogClientConfig struct {
		LoggerName string `mapstructure:"loggername"`
	}
	LogServerConfig struct {
		Address    string `mapstructure:"address"`
		LoggerName string `mapstructure:"loggername"`
	}
)

func (c *CollectorConfig) Complete() {
	//c.RPC.Complete()
	//c.Log.Complete()
}

func (c *CollectorConfig) Validate() bool {
	//return c.RPC.Validate() && c.Log.Validate()
	return true
}

func (c *RPCConfig) Complete() {
}

func (c *RPCConfig) Validate() bool {
	return true
}

//func (c *LogConfig) Complete() {
//}
//
//func (c *LogConfig) Validate() bool {
//	return true
//}

func NewCollectorConfig(confPath string, configType string) (*CollectorConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName(confPath)
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType(configType)
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(CollectorConfig)
	err := viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	conf.Complete()
	if ok := conf.Validate(); !ok {
		return nil, errorx.ErrConfigInvalid
	}
	return conf, nil
}
