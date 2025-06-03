package main

import (
	"net/http"
	"rest-api/common/middleware"
	"rest-api/features/post"
	"rest-api/features/secret"

	"github.com/go-playground/validator/v10"
)

func main() {
	r := http.NewServeMux()

	middlewares := []middleware.Middleware{
		middleware.LogMiddleware,
		middleware.JsonMiddleware,
	}

	db := post.NewPostRepository()
	v := validator.New(validator.WithRequiredStructEnabled())

	postHandlers := post.PostHandlers{
		Repository: db,
		Validator:  v,
	}
	r.HandleFunc("GET /posts", postHandlers.Index)
	r.HandleFunc("GET /posts/{id}", postHandlers.Show)
	r.HandleFunc("POST /posts", postHandlers.Create)
	r.HandleFunc("DELETE /posts/{id}", postHandlers.Delete)

	secretHandlers := secret.SecretHandlers{}
	jwtMiddleware := middleware.NewChain(middleware.JwtMiddleware)
	r.HandleFunc("GET /secrets", jwtMiddleware.Handle(secretHandlers.Index))
	r.HandleFunc("GET /secrets/{wildcard}", jwtMiddleware.Handle(secretHandlers.Show))

	globalMiddleware := middleware.NewChain(middlewares...)
	http.ListenAndServe(":3000", globalMiddleware.Apply(r))
}
