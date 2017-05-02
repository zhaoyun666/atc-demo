# ATC

ATC 是一个快速开发GO应用程序的开源框架，支持RESTful API 及 Thrift RPC的框架.可根据自身业务逻辑选择性的卸载中间件的功能，均支持平滑退出。

DEMO案例

获取ATC框架源码 [atc.wiki](https://github.com/adolphlxm/atc)

## ATC项目结构
<pre>
├── conf
│   ├── app.ini
│   └── error.ini
├── front
│   └── HTML...
├── bin
├── src
│   ├── httprouter
│   │    ├── V1
│   │    └── router.go
│   └── thriftrpc
│         ├── idl
│         ├── gen
│         ├── ...(.go)
│         └── router.go
└── atc.go
</pre>

## 快速入门

> 该DEMO以 users 模块 写了一个简单案例仅供参考，目录结构也仅供参考。

### RESTful API

* httprouter 为 HTTP请求路由
    - V1 版本1
    - V2 版本2
    - router.go 为了方便管理路由, 统一把注入路由放到根目录`router.go`文件内了
    
* DEMO
    - [GET] http://127.0.0.1:9001/V1/users/login?flag=true&id=1
        <pre>
        {
          "ID": "1",
          "data": {
            "regtime": "2017-04-28",
            "username": "ATC-V!"
          }
        }
        </pre>
    - [POST] http://127.0.0.1:9001/V1/users/login?username=atc&password=123456
        
        返回JSON
    - .....
    - [GET] .../V2/users/login?id=1
    
    可以自行测试下结果！

### model 
> DB层相关代码逻辑

### service
> 逻辑层相关代码逻辑

### ThriftRPC

* thriftrpc 为thrift通信RPC
    - idl IDL文件存放
    - gen 通过thrift生成的go文件存放
    - login.go ... 这里写相关通信调用封装
    - router.go 为了方便管理路由，统一把多路复用注册放到根目录`router.go`文件内了
    
注：如果发生idl结构变化，需要重新部署项目，建议可以生成新的idl来避免新老版本idl兼容问题。

* DEMO
> thriftrpc/idl/micro.thrift 是一个简单的IDL文件
>
> thriftrpc/gen 为thrift生成的go文件

```thrift 
namespace go micro
namespace php micro


struct Article{
 1: i32 id, 
 2: string title,
 3: string content,
 4: string author,
}
 
const map<string,string> MAPCONSTANT = {'hello':'world', 'goodnight':'moon'}
 
service microThrift {        
        list<string> CallBack(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
        void put(1: Article newArticle),
}
```
login.go 实现了 `CallBack()` 和 `put()` 两个方法
 
通过PHP/Python/Go 等自行实现客户端进行测试

**注意：thrift client 使用的transportT 和 protocolT 需要和 框架app.ini 配置内一致！！！**
