package main

import (
	_ "github.com/1uvu/bitlog/collector"
	_ "github.com/1uvu/bitlog/parser"
	_ "github.com/1uvu/bitlog/pkg/common"
	_ "github.com/1uvu/bitlog/pkg/config"
	_ "github.com/1uvu/bitlog/pkg/errorx"
	_ "github.com/1uvu/bitlog/pkg/utils"
	_ "github.com/1uvu/bitlog/storage"
)

// TODO 添加各种启动参数，启动不同的服务
//  - 日志接受 http 服务
//  - 等等

func main() {

}
