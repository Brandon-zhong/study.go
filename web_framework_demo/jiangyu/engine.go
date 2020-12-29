package jiangyu

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) Get(path string, handler HandlerFunc) {
	e.router.addRoute("GET", path, handler)
}
func (e *Engine) Post(path string, handler HandlerFunc) {
	e.router.addRoute("POST", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(newContext(w, r))
}
