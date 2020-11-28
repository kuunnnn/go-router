package grouter

import (
	"fmt"
	"go-router/tool"
	"go-router/trie"
	"net/http"
	"os"
)

type RouterContext struct {
	req    *http.Request
	res    http.ResponseWriter
	params map[string]string
}

func (ctx *RouterContext) Send() {

}

type Handler func(ctx *RouterContext)

type Methods = int

const (
	GET Methods = iota
	POST
	PUT
	DELETE
	HEAD
	OPTIONS
	PATCH
)

type Routes interface {
	Get(string, ...Handler)
	Post(string, ...Handler)
	Put(string, ...Handler)
	Delete(string, ...Handler)
	Patch(string, ...Handler)
	Options(string, ...Handler)
	Head(string, ...Handler)

	// 暴露一个文件到外部
	StaticFile(path string, source string)
	// 暴露一个目录内的文件到外部
	Static(path string, source string)
	// 暴露一个目录内的文件到外部
	// 可以获取目录列表
	StaticFS(path string, source string)

	Group(prefix string) *groupRouter
	Mount(port string)
}

type tepHandlerList struct {
	path     string
	method   Methods
	handlers []Handler
}

type router struct {
	prefix      string
	root        bool
	trees       map[string]*trie.Trie
	groupRouter []*groupRouter
	hs          []*tepHandlerList
}

type groupRouter struct {
	router
}

type service struct {
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// 调用子路由的 mount 方法构建出一颗 trie 树
func (r *router) Mount(port string) {
	if !r.root {
		return
	}
	queue := tool.NewQueue(100)
	type tmp struct {
		router *groupRouter
		prefix string
	}
	for _, routes := range r.groupRouter {
		queue.Push(&tmp{router: routes, prefix: r.prefix})
	}
	for queue.Size > 0 {
		e := queue.Shift()
		for _, h := range e.(tmp).router.hs {
			r.trees[string(h.method)].Insert(e.(tmp).prefix+h.path, h.handlers)
		}
		for _, routes := range e.(tmp).router.groupRouter {
			queue.Push(&tmp{router: routes, prefix: e.(tmp).prefix})
		}
	}
	service := &service{}
	service.ServeHTTP = func(w http.ResponseWriter, req *http.Request) {
		handler, params := r.trees[req.Method].Match(req.URL.Path)
		handler.(Handler)(&RouterContext{
			req:    req,
			res:    w,
			params: params,
		})
	}
	err := http.ListenAndServe(port, service)
	if err != nil {
		fmt.Printf("setup service[%s] error %+v\n", port, err)
		os.Exit(-1)
	}
}

func New() Routes {
	return &router{
		prefix:      "/",
		root:        true,
		trees:       map[string]*trie.Trie{},
		hs:          make([]*tepHandlerList, 0, 10),
		groupRouter: []*groupRouter{},
	}
}

// 创建一个子路由
func (r *router) Group(prefix string) *groupRouter {
	route := &groupRouter{
		router{
			prefix:      r.prefix + prefix,
			root:        false,
			trees:       map[string]*trie.Trie{},
			hs:          make([]*tepHandlerList, 0, 10),
			groupRouter: []*groupRouter{},
		},
	}
	r.groupRouter = append(r.groupRouter, route)
	return route
}

func (r *router) addRoute(methods Methods, path string, handler ...Handler) {
	r.hs = append(r.hs, &tepHandlerList{path: path, handlers: handler, method: methods})
}
func (r *router) compose(handlers ...Handler) Handler {
	return func(ctx *RouterContext) {
		for _, handler := range handlers {
			handler(ctx)
		}
	}
}

func (r *router) Get(path string, handler ...Handler) {
	r.addRoute(GET, path, handler...)
}
func (r *router) Post(path string, handler ...Handler) {
	r.addRoute(POST, path, handler...)
}
func (r *router) Put(path string, handler ...Handler) {
	r.addRoute(PUT, path, handler...)
}
func (r *router) Delete(path string, handler ...Handler) {
	r.addRoute(DELETE, path, handler...)
}
func (r *router) Patch(path string, handler ...Handler) {
	r.addRoute(PATCH, path, handler...)
}
func (r *router) Options(path string, handler ...Handler) {
	r.addRoute(OPTIONS, path, handler...)
}
func (r *router) Head(path string, handler ...Handler) {
	r.addRoute(HEAD, path, handler...)
}

// 指定一个文件访问路径映射
func (r *router) StaticFile(path string, source string) {
	r.addRoute(GET, path, func(ctx *RouterContext) {

	})
}

// 可以访问目录下的所有文件, 但是不能访问目录
func (r *router) Static(path string, source string) {
	r.addRoute(GET, path, func(ctx *RouterContext) {

	})
}

// 包含 Static 功能并在访问的是一个目录的时候会展开目录结构
func (r *router) StaticFS(path string, source string) {
	r.addRoute(GET, path, func(ctx *RouterContext) {

	})
}
