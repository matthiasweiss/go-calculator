package main

import (
	"net/http"
	"rest-api/common/middleware"
	"rest-api/features/post"
)

func main() {
	r := http.NewServeMux()

	routers := map[string]http.Handler{
		"/posts": post.NewPostRouter(),
	}

	for prefix, subRouter := range routers {
		r.Handle(prefix, http.StripPrefix(prefix, subRouter))
	}

	middlewares := []middleware.Middleware{
		middleware.LogMiddleware,
		middleware.JsonMiddleware,
	}

	middlewareChain := middleware.NewChain(middlewares...)

	http.ListenAndServe(":3000", middlewareChain.Apply(r))
}
