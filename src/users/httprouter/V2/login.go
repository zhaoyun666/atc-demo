package userV2

import (
	"github.com/adolphlxm/atc"
)

type LoginHandler struct {
	atc.Handler
}

//200 OK - [GET]：服务器成功返回用户请求的数据，该操作是幂等的（Idempotent）。
//201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功。
//202 Accepted - [*]：表示一个请求已经进入后台排队（异步任务）
//204 NO CONTENT - [DELETE]：用户删除数据成功。
//400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
//401 Unauthorized - [*]：表示用户没有权限（令牌、用户名、密码错误）。
//403 Forbidden - [*] 表示用户得到授权（与401错误相对），但是访问是被禁止的。
//404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
//406 Not Acceptable - [GET]：用户请求的格式不可得（比如用户请求JSON格式，但是只有XML格式）。
//410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的。
//422 Unprocesable entity - [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误。
//500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。

// 从服务器取出资源(一项或多项)
func (this *LoginHandler) Get() {
	// 已登录
	if this.Ctx.Query("flag") == "true" {
		loginData := map[string]interface{}{
			"username": "ATC-V2",
			"regtime":  "2017-04-28",
		}
		this.Ctx.SetData("data", loginData)
		this.Ctx.SetData("ID", this.Ctx.Query("id"))
		this.JSON()
		return
	}

	// 未登录
	// 自定义错误提示内容
	//this.Error406(-1).Message("没有权限查看").JSON()
	// error.ini错误码匹配的提示内容
	this.Error406(-1).JSON()
}

// 从服务器新建一个资源
func (this *LoginHandler) Post() {
	username := this.Ctx.Query("username")
	password := this.Ctx.Query("password")
	if username == "" || password == "" {
		this.Error406(10001)
		return
	}
	loginData := map[string]interface{}{
		"username": username,
		"password": password,
		"status":   "Insert V2 POST success.",
	}
	this.Ctx.SetData("result", loginData)
	this.JSON()
}

// 在服务器更新资源(客户端提供改变后的完整资源)
func (this *LoginHandler) Put() {
	username := this.Ctx.Query("username")
	password := this.Ctx.Query("password")
	if username == "" || password == "" {
		this.Error406(10001)
		return
	}
	loginData := map[string]interface{}{
		"username": username,
		"password": password,
		"status":   "Update V2 PUT success.",
	}
	this.Ctx.SetData("result", loginData)
	this.JSON()
}

// 在服务器更新资源（客户端提供改变的属性）
func (this *LoginHandler) Patch() {
	id := this.Ctx.Query("id")
	if id == "" {
		this.Error406(10001)
		return
	}
	loginData := map[string]interface{}{
		"id":     id,
		"status": "Update V2 PATCH success.",
	}
	this.Ctx.SetData("result", loginData)
	this.JSON()
}

// 从服务器删除资源
func (this *LoginHandler) Delete() {
	id := this.Ctx.Query("id")
	if id == "" {
		this.Error406(10001)
		return
	}
	loginData := map[string]interface{}{
		"id":     id,
		"status": "Delete V2 DELETE success.",
	}
	this.Ctx.SetData("result", loginData)
	this.JSON()
}
