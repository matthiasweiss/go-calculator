package routes

import (
	"log"
	"net/http"
	"rest-api/common/middleware"
	"rest-api/features/post"
	"rest-api/features/secret"
)

type Handlers struct {
	Post   *post.PostHandlers
	Secret *secret.SecretHandlers
}

func SetupRoutes(mux *http.ServeMux, h Handlers, l *log.Logger) {
	mux.HandleFunc("GET /posts", h.Post.Index)
	mux.HandleFunc("GET /posts/{id}", h.Post.Show)
	mux.HandleFunc("POST /posts", h.Post.Create)
	mux.HandleFunc("DELETE /posts/{id}", h.Post.Delete)

	jwtMiddleware := middleware.JwtMiddleware(l)
	jwtMiddlewareChain := middleware.NewChain(jwtMiddleware)
	mux.HandleFunc("GET /secrets", jwtMiddlewareChain.Handle(h.Secret.Index))
	mux.HandleFunc("GET /secrets/{wildcard}", jwtMiddlewareChain.Handle(h.Secret.Show))
}
