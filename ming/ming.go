package ming

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router.addRoute(method, pattern, handlerFunc)
}
func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}
func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}
func (engine *Engine) PUT(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("PUT", pattern, handlerFunc)
}
func (engine *Engine) DELETE(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("DELETE", pattern, handlerFunc)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
