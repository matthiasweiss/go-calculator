package main

import (
	"encoding/json"
	"net/http"
	"rest-api/data"
	"rest-api/database"
	"rest-api/validation"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func main() {
	r := http.NewServeMux()

	db := database.NewDatabase()

	v := validator.New(validator.WithRequiredStructEnabled())

	r.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		posts := db.Index()

		err := json.NewEncoder(w).Encode(posts)

		if err != nil {
			http.Error(w, "Could not encode posts", http.StatusBadRequest)
			return
		}
	})

	r.HandleFunc("POST /posts", func(w http.ResponseWriter, r *http.Request) {
		var postData data.PostData
		json.NewDecoder(r.Body).Decode(&postData)

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

	r.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		param := r.PathValue("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		post, err := db.Show(id)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(post)

		if err != nil {
			http.Error(w, "Could not encode posts", http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(":3000", r)
}
