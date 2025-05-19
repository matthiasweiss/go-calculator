package handlers

import (
	"encoding/json"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Message string `json:"message"`
	}

	result := Result{
		Message: "Hello World",
	}

	json.NewEncoder(w).Encode(result)
}
