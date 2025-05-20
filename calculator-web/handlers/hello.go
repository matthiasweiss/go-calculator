package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Message string `json:"message"`
	}

	result := Result{
		Message: "Hello World",
	}

	fmt.Println(r.URL.Query())

	json.NewEncoder(w).Encode(result)
}

func Wildcard(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Param int `json:"param"`
	}

	param, err := strconv.Atoi(r.PathValue("param"))

	if err != nil {
		http.Error(w, "{param} must be an int", http.StatusBadRequest)
		return
	}

	result := Result{
		Param: param,
	}

	json.NewEncoder(w).Encode(result)
}
