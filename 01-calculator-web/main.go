package main

import (
	"calculator/handlers"
	"calculator/router"
	"fmt"
	"net/http"
)

/*
Calculator API, spec taken from https://github.com/dreamsofcode-io/goprojects/blob/main/02-backend-api/README.md
*/
func main() {
	router := router.New()

	router.Get("/", handlers.Hello)
	router.Get("/hello/{param}", handlers.Wildcard)
	router.Post("/add", handlers.Add)
	router.Post("/subtract", handlers.Subtract)
	router.Post("/multiply", handlers.Multiply)
	router.Post("/divide", handlers.Divide)
	router.Post("/sum", handlers.Sum)

	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)

	// keep in here for now, maybe I'll implement this completely
	// http.ListenAndServe(":8080", router)
}
