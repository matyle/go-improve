# GIN 框架

## 开始使用
- 创建 example 文件夹
- go mod init
- 创建 example.go 

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	root := r.Group("/test")

	root.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":18089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

## 从 example 开始
`gin.go`文件

### 设置默认内存大小
```go
const defaultMultipartMemory = 32 << 20 // 32 M
```

### 函数作为一等公民：当作类型，变量，值使用
- 放在切片中
- 类型方法
- 作为参数传递
```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc slice.
type HandlersChain []HandlerFunc

// Last returns the last handler in the chain. i.e. the last handler is the main one.
func (c HandlersChain) Last() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}
```
## Gin.go中的数据结构
### 结构体 
#### 路由信息
```go
// RouteInfo represents a request route's specification which contains method and path and its handler.
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc //函数类型
}

// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}
```
#### Engine
```go
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
1. Engine 中含有RouterGroup匿名类型，而RouteGroup有 Engine 指针?
2. 

type Engine strcut{
	RouterGroup

}

```

## 函数


#### HandlerFunc
```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context) //无返回值，带*Context参数

// HandlersChain defines a HandlerFunc slice.
type HandlersChain []HandlerFunc

```


#### New 初始化 Engine
New returns a new blank Engine instance without any middleware attached. (不带任何中间件)
初始化 Engine 结构体
1. 首先初始化 Engine 中多数成员变量
2. 初始化 Engine 中 RouterGroup中的 内部成员 engine 变量 (有什么作用呢？)
3. 初始化 Engine 中 sync.Pool的 pool 成员(存储临时对象)

```go
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		RemoteIPHeaders:        []string{"X-Forwarded-For", "X-Real-IP"},
		TrustedPlatform:        defaultPlatform,
		UseRawPath:             false,
		RemoveExtraSlash:       false,
		UnescapePathValues:     true,
		MaxMultipartMemory:     defaultMultipartMemory,
		trees:                  make(methodTrees, 0, 9),
		delims:                 render.Delims{Left: "{{", Right: "}}"},
		secureJSONPrefix:       "while(1);",
		trustedProxies:         []string{"0.0.0.0/0", "::/0"},
		trustedCIDRs:           defaultTrustedCIDRs,
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() any {
		return engine.allocateContext()
	}
	return engine
}
```


#### Default
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
调用 New 方法
调用 Use 方法，添加中间件
```go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```
## 重要方法Method

```go
//添加中间件
(*Engine).Use()

//添加路由
(*Engine).addRoute()

//Run 系列
(*Engine).Run()

//HTTP系列
(*Engine).ServeHTTP()
```
####  中间件Middleware（Use 方法）
Use 方法为路由添加一个全局中间件，包括每一个单独请求的所有的处理函数链
```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}
```
这里主要调用了 RouterGroup的 Use 方法[[RouterGroup.md]]
####  (\*Engine).ServeHTTP()
将请求转到gin中, 使用gin的相关函数处理request请求

