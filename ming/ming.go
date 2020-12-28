package ming

import (
	"log"
	"net/http"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type HandlerFunc func(ctx *Context)
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}
func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("GET", pattern, handlerFunc)
}
func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("POST", pattern, handlerFunc)
}
func (group *RouterGroup) PUT(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("PUT", pattern, handlerFunc)
}
func (group *RouterGroup) DELETE(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("DELETE", pattern, handlerFunc)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
