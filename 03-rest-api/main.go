package main

import (
	"net/http"
	"rest-api/common/middleware"
	"rest-api/features/post"
	"rest-api/features/secret"
)

func main() {
	r := http.NewServeMux()

	middlewares := []middleware.Middleware{
		middleware.LogMiddleware,
		middleware.JsonMiddleware,
	}

	chain := middleware.NewChain(middlewares...)

	post.RegisterRoutes(r)
	secret.RegisterRoutes(r)

	http.ListenAndServe(":3000", chain.Apply(r))
}
