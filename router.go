package go_router

// 路由分静态路由和动态路由
type Router struct {
	hRouter *hRouter
	tRouter *Trie
}
type Handler func()

type IRouter interface {
	AddRoute()
	MatchRoute() Handler
}

// 添加路由:
//		1. 判断是动态路由还是静态路由
//		2. 静态路由用 hash 表存储, 动态路由使用 trie 树存储
func (r *Router) AddRoute() {

}

// 查找路由:
// 		1. 先用静态路由匹配
// 		2. 如果静态路由匹配失败再使用进行动态路由匹配
// 	    3. 如果也失败则 404
func (r *Router) MatchRoute() Handler {
	return nil
}
