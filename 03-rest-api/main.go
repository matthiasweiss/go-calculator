package main

import (
	"net/http"
	"rest-api/middleware"
	"rest-api/post"
)

func main() {

	r := http.NewServeMux()

	routers := map[string]http.Handler{
		"/posts": post.NewPostRouter(),
	}

	for prefix, subRouter := range routers {
		r.Handle(prefix, http.StripPrefix(prefix, subRouter))
	}

	chain := middleware.NewMiddlewareChain(middleware.Log, middleware.Json)
	http.ListenAndServe(":3000", chain.Apply(r))
}
