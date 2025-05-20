package handlers

import (
	"encoding/json"
	"net/http"
)

type Calculation struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Result struct {
	Result float64 `json:"result"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var calculation Calculation
	json.NewDecoder(r.Body).Decode(&calculation)

	result := Result{
		Result: calculation.Number1 + calculation.Number2,
	}

	json.NewEncoder(w).Encode(result)
}

func Subtract(w http.ResponseWriter, r *http.Request) {
	var calculation Calculation
	json.NewDecoder(r.Body).Decode(&calculation)

	result := Result{
		Result: calculation.Number1 - calculation.Number2,
	}

	json.NewEncoder(w).Encode(result)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	var calculation Calculation
	json.NewDecoder(r.Body).Decode(&calculation)

	result := Result{
		Result: calculation.Number1 * calculation.Number2,
	}

	json.NewEncoder(w).Encode(result)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var calculation Calculation
	json.NewDecoder(r.Body).Decode(&calculation)

	result := Result{
		Result: calculation.Number1 / calculation.Number2,
	}

	json.NewEncoder(w).Encode(result)
}

func Sum(w http.ResponseWriter, r *http.Request) {
	type SumCalculation struct {
		Items []float64 `json:"items"`
	}
	var calculation SumCalculation
	json.NewDecoder(r.Body).Decode(&calculation)

	result := Result{
		Result: sum(calculation.Items),
	}

	json.NewEncoder(w).Encode(result)
}

func sum(items []float64) float64 {
	sum := 0.0

	for _, item := range items {
		sum += item
	}

	return sum
}
