package main

import (
	"calculator/handlers"
	"calculator/router"
	"net/http"
)

/*
Calculator API, spec taken from https://github.com/dreamsofcode-io/goprojects/blob/main/02-backend-api/README.md
*/
func main() {
	router := router.New()

	router.Post("/add", handlers.Add)
	router.Post("/subtract", handlers.Subtract)
	router.Post("/multiply", handlers.Multiply)
	router.Post("/divide", handlers.Divide)
	router.Post("/sum", handlers.Sum)

	http.ListenAndServe(":8080", router)
}
