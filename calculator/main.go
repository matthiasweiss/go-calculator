package main

import (
	"calculator/router"
	"encoding/json"
	"fmt"
	"net/http"
)

type Calculation struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Result struct {
	Result float64 `json:"result"`
}

/*
Calculator API, spec taken from https://github.com/dreamsofcode-io/goprojects/blob/main/02-backend-api/README.md
*/
func main() {
	router := router.New()

	router.Post("/add", func(w http.ResponseWriter, r *http.Request) {
		var calculation Calculation
		json.NewDecoder(r.Body).Decode(&calculation)

		result := Result{
			Result: calculation.Number1 + calculation.Number2,
		}

		json.NewEncoder(w).Encode(result)
	})

	router.Post("/subtract", func(w http.ResponseWriter, r *http.Request) {
		var calculation Calculation
		json.NewDecoder(r.Body).Decode(&calculation)

		result := Result{
			Result: calculation.Number1 - calculation.Number2,
		}

		json.NewEncoder(w).Encode(result)
	})

	router.Post("/multiply", func(w http.ResponseWriter, r *http.Request) {
		var calculation Calculation
		json.NewDecoder(r.Body).Decode(&calculation)

		result := Result{
			Result: calculation.Number1 * calculation.Number2,
		}

		json.NewEncoder(w).Encode(result)
	})

	router.Post("/divide", func(w http.ResponseWriter, r *http.Request) {
		var calculation Calculation
		json.NewDecoder(r.Body).Decode(&calculation)

		result := Result{
			Result: calculation.Number1 / calculation.Number2,
		}

		json.NewEncoder(w).Encode(result)
	})

	router.Post("/sum", func(w http.ResponseWriter, r *http.Request) {
		type SumCalculation struct {
			Items []float64 `json:"items"`
		}
		var calculation SumCalculation
		json.NewDecoder(r.Body).Decode(&calculation)

		fmt.Println(calculation.Items)

		result := Result{
			Result: sum(calculation.Items),
		}

		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}

func sum(items []float64) float64 {
	sum := 0.0

	for _, item := range items {
		fmt.Println(item)
		sum += item
	}

	return sum
}
