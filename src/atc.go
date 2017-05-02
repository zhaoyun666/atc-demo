package main

import (
	"github.com/adolphlxm/atc"
	_ "github.com/adolphlxm/atc-demo/src/users/httprouter"
	_ "github.com/adolphlxm/atc-demo/src/users/thriftrpc"
)

func main() {
	// 根据配置文件注入依赖中间件
	// 目前支持：HTTP/Websoeckt、Thrift、ORM(xorm、其它待开发)
	atc.Run()
}
