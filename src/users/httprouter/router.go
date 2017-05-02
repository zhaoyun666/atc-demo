package httprouter

import (
	"github.com/adolphlxm/atc"
	"github.com/adolphlxm/atc-demo/src/users/httprouter/V1"
)

func init() {
	v1 := atc.Route.Group("V1")
	{
		// V1版本过滤器, 根据路由规则加载。
		// 支持三种过滤器：
		//      1. EFORE_ROUTE                    //匹配路由之前
		//      2. BEFORE_HANDLER                 //匹配到路由后,执行Handler之前
		//      3. AFTER                          //执行完所有逻辑后
		v1.AddFilter(atc.BEFORE_ROUTE, "users.*", userV1.AfterLogin)
		v1.AddRouter("users", &userV1.LoginHandler{})
	}

	v2 := atc.Route.Group("V2")
	{
		v2.AddRouter("users", &userV1.LoginHandler{})
	}

}
