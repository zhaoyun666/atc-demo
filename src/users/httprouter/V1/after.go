package userV1

import (
	"github.com/adolphlxm/atc/context"
)

// 登录过滤器
// 可以通过自定义过滤器来实现登录状态、权限检查等功能
func AfterLogin(ctx *context.Context) {
	// 错误输出
	//error := atc.NewError(ctx)
	//error.Code(401,10000).JSON()
}
