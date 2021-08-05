package main

import (
	"net/http"
	"strings"
	"sync"
)

type HandlerFunc func(*Context) error

type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
	pool   *sync.Pool
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{Engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.AddRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodGet, pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPost, pattern, handler)
}
func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPut, pattern, handler)
}
func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodDelete, pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) {
	err := http.ListenAndServe(addr, engine)
	if err != nil {
		return
	}
}

// Run defines the method to start a https server
func (engine *Engine) RunTLS(addr, certFile, keyFile string) {
	err := http.ListenAndServeTLS(addr, certFile, keyFile, engine)
	if err != nil {
		return
	}
	return
}

func (engine *Engine) middlewares(path string) (middlewares []HandlerFunc) {
	for _, group := range engine.groups {
		if strings.HasPrefix(path, group.Prefix) {
			middlewares = append(middlewares, group.Middlewares...)
		}
	}
	return
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	c := newContext(w, req)
	c.handlers = engine.middlewares(req.URL.Path)
	engine.router.Handle(c)
	//engine.pool.Put(c)
}
func (engine *Engine) allocateContext() *Context {
	v := make(map[string]string, 0)
	return &Context{engine: engine, Params: v}
}
