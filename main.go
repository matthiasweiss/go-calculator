package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Result int `json:"result"`
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		result := Result{
			Result: 12,
		}

		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("POST /decode", func(w http.ResponseWriter, r *http.Request) {
		var result Result
		json.NewDecoder(r.Body).Decode(&result)

		fmt.Fprintf(w, "Result is: %d", result.Result)
	})

	http.ListenAndServe(":8080", nil)
}
