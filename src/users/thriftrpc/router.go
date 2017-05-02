package thriftrpc

import (
	"github.com/adolphlxm/atc"
	"github.com/adolphlxm/atc-demo/src/users/thriftrpc/gen/micro"
)

func init() {
	// Thrift RPC 多路复用路由注册
	processor := micro.NewMicroThriftProcessor(&MicroHandler{})
	atc.ThriftRPC.RegisterProcessor("user", processor)
}
