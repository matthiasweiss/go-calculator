package secret

import (
	"net/http"
	"rest-api/common/middleware"
)

type secret struct {
	Message string `json:"message"`
}

func RegisterRoutes(r *http.ServeMux) {
	handlers := SecretHandlers{}

	middlewares := []middleware.Middleware{
		middleware.JwtMiddleware,
	}

	chain := middleware.NewChain(middlewares...)

	r.HandleFunc("GET /secrets", chain.Handle(handlers.Index))
	r.HandleFunc("GET /secrets/{wildcard}", chain.Handle(handlers.Show))
}
