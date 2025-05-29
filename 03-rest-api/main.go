package main

import (
	"net/http"
	"rest-api/database"
	"rest-api/handlers"
	"rest-api/middleware"

	"github.com/go-playground/validator/v10"
)

func main() {
	r := http.NewServeMux()

	db := database.NewDatabase()
	v := validator.New(validator.WithRequiredStructEnabled())

	postHandlers := handlers.PostHandlers{
		Database:  db,
		Validator: v,
	}

	r.HandleFunc("GET /posts", postHandlers.Index)
	r.HandleFunc("GET /posts/{id}", postHandlers.Show)
	r.HandleFunc("POST /posts", postHandlers.Create)
	r.HandleFunc("DELETE /posts/{id}", postHandlers.Delete)

	chain := middleware.NewMiddlewareChain(middleware.Log, middleware.Json)

	http.ListenAndServe(":3000", chain.Apply(r))
}
