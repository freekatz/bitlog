package config

import (
	"time"

	"github.com/1uvu/bitlog/pkg/errorx"

	"github.com/spf13/viper"
)

type (
	CollectorConfig struct {
		/*
			client config
		*/
		RPC   *RPConfig    `mapstructure:"rpc"`
		Shell *ShellConfig `mapstructure:"shell"`

		/*
			output config
		*/
		Output *OutputConfig `mapstructure:"output"`

		///*
		//	mode configs
		//*/
		//Mode RunMode `mapstructure:"mode"`
		//// for block mode
		//BlockCount int64 `mapstructure:"blockCount"`
		//// for duration mode
		//Period   string `mapstructure:"period"`
		//Duration string `mapstructure:"duration"`

		///*
		//	type configs
		//*/
		//Type NetType `mapstructure:"type"`
	}
	RPConfig struct {
		Address  string `mapstructure:"address"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
	ShellConfig struct {
	}
	OutputConfig struct {
		Path string `mapstructure:"path"`
	}
	//RunMode string
	//NetType string
)

//const (
//	Once     RunMode = "once"
//	Block    RunMode = "block"
//	Duration RunMode = "duration"
//	Infinity RunMode = "infinity"
//)
//
//const (
//	Testnet NetType = "testnet"
//	Simnet  NetType = "simnet"
//	Mainnet NetType = "mainnet"
//	Signet  NetType = "signet"
//)
//
//var (
//	defaultLogPath    = "./logs"
//	defaultRunMode    = Infinity
//	defaultBlockCount = 6
//	defaultPeriod     = "5m"
//	defaultDuration   = "1h"
//)

func (c *CollectorConfig) Complete() {
}

func (c *CollectorConfig) Validate() bool {
	return true
}

func (c *RPConfig) Complete() {
}

func (c *RPConfig) Validate() bool {
	return true
}

func ValidateDuration(duration string) bool {
	d, err := time.ParseDuration(duration)
	return err == nil && d > 0
}

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
