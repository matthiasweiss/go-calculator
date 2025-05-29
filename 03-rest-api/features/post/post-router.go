package post

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

func NewPostRouter() *http.ServeMux {
	db := NewPostRepository()
	v := validator.New(validator.WithRequiredStructEnabled())

	postHandlers := PostHandlers{
		Database:  db,
		Validator: v,
	}

	r := http.NewServeMux()

	r.HandleFunc("GET /posts", postHandlers.Index)
	r.HandleFunc("GET /posts/{id}", postHandlers.Show)
	r.HandleFunc("POST /posts", postHandlers.Create)
	r.HandleFunc("DELETE /posts/{id}", postHandlers.Delete)

	return r
}
