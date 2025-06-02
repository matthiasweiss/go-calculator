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

	chain := middleware.NewChain(middleware.JwtMiddleware)

	r.HandleFunc("GET /secrets", chain.Handle(handlers.Index))
	r.HandleFunc("GET /secrets/{wildcard}", chain.Handle(handlers.Show))
}
