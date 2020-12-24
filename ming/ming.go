package ming

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
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
	key := r.Method + "-" + r.URL.Path
	if handlerFunc, ok := engine.router[key]; ok {
		handlerFunc(w, r)
	} else {
		fmt.Fprintf(w, "404:%q\n", r.URL.Path)
	}
}
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
