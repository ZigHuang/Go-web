package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine 实现 ServeHTTP 接口
type (
	RouterGroup struct {
		// 分组前缀 例如: /api
		prefix string
		// 中间件支持
		middlewares []HandlerFunc
		// 支持嵌套
		parent *RouterGroup
		// 所有分组共享一个 Engine 实例
		engine *Engine
	}
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 创建一个新的路由分组
// Ps: 上面提到 多个分组共用一个 engine 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	// 一个新的分组
	newGroup := &RouterGroup{
		// 前缀叠加 例如：/api + /xxx/delete
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 下面关于路由/请求的方法由原来的 Engine 处理变为交给 RouterGroup 处理

// addRoute 添加路由信息
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	// 路由前缀拼装
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 解析 GET 请求
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 解析 POST 请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run 定义一个开启 http server 的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Use 为某一路由分组添加中间件支持
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 实现 ServeHTTP 接口
// PS：实现了接口方法的 struct 都可以强制转换为接口类型
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		// 判断该请求适用于那些中间件
		// 对应的一个 RouterGroup 下可以有多个中间件支持
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	// 将中间件列表赋值给c
	c.handlers = middlewares
	engine.router.handle(c)
}
