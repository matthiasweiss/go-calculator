package main

import (
	"encoding/json"
	"net/http"
	"rest-api/data"
	"rest-api/database"
	"rest-api/validation"

	"github.com/go-playground/validator/v10"
)

func main() {
	r := http.NewServeMux()

	db := database.NewDatabase()

	v := validator.New(validator.WithRequiredStructEnabled())

	r.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		posts := db.List()

		err := json.NewEncoder(w).Encode(posts)

		if err != nil {
			http.Error(w, "Could not encode posts", http.StatusBadRequest)
			return
		}
	})

	r.HandleFunc("POST /posts", func(w http.ResponseWriter, r *http.Request) {
		var postData data.PostData
		json.NewDecoder(r.Body).Decode(&postData)
		w.Header().Set("Content-Type", "application/json")

		err := v.Struct(postData)

		if err != nil {
			errorResponse := validation.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		post := db.Create(postData)

		err = json.NewEncoder(w).Encode(post)

		if err != nil {
			http.Error(w, "Could not encode posts", http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(":3000", r)
}
