package main

import grouter "go-router"

func handler(ctx *grouter.RouterContext) {

}

func main() {
	rootRouter := grouter.New()
	userRouter := rootRouter.Group("/user")
	orderRouter := rootRouter.Group("/order")

	userRouter.Get("/", handler)
	userRouter.Get("/info", handler)
	userRouter.Put("/info", handler)
	userRouter.Post("/login", handler)

	userRouter.Get("/", handler)
	orderRouter.Get("/list", handler)
	orderRouter.Get("/info", handler)
	orderRouter.Put("/info", handler)

	rootRouter.StaticFS("/static", "/")
	rootRouter.Mount(":8080")
}
