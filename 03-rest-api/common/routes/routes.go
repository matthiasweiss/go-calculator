package routes

import (
	"net/http"
	"rest-api/common/middleware"
	"rest-api/features/post"
	"rest-api/features/secret"

	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	Post   *post.PostHandlers
	Secret *secret.SecretHandlers
}

func SetupRoutes(r *http.ServeMux, h Handlers) {
	r.HandleFunc("GET /posts", h.Post.Index)
	r.HandleFunc("GET /posts/{id}", h.Post.Show)
	r.HandleFunc("POST /posts", h.Post.Create)
	r.HandleFunc("DELETE /posts/{id}", h.Post.Delete)

	jwtMiddleware := middleware.NewChain(middleware.JwtMiddleware)
	r.HandleFunc("GET /secrets", jwtMiddleware.Handle(h.Secret.Index))
	r.HandleFunc("GET /secrets/{wildcard}", jwtMiddleware.Handle(h.Secret.Show))
}
