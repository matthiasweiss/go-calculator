package secret

import (
	"net/http"
)

type secret struct {
	Message string `json:"message"`
}

func RegisterRoutes(r *http.ServeMux) {
	handlers := SecretHandlers{}

	r.HandleFunc("GET /secrets", handlers.Index)
	r.HandleFunc("GET /secrets/{wildcard}", handlers.Show)
}
