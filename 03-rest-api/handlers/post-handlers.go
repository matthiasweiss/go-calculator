package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/data"
	"rest-api/database"
	"rest-api/validation"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type PostHandlers struct {
	Database  *database.Database
	Validator *validator.Validate
}

func (h *PostHandlers) Index(w http.ResponseWriter, r *http.Request) {
	posts := h.Database.Index()

	err := json.NewEncoder(w).Encode(posts)

	if err != nil {
		http.Error(w, "Could not encode posts", http.StatusBadRequest)
		return
	}
}

func (h *PostHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var postData data.PostData
	json.NewDecoder(r.Body).Decode(&postData)

	err := h.Validator.Struct(postData)

	if err != nil {
		errorResponse := validation.NewErrorResponse(err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	post := h.Database.Create(postData)

	err = json.NewEncoder(w).Encode(post)

	if err != nil {
		http.Error(w, "Could not encode posts", http.StatusBadRequest)
		return
	}
}

func (h *PostHandlers) Show(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	post, err := h.Database.Show(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(post)

	if err != nil {
		http.Error(w, "Could not encode posts", http.StatusBadRequest)
		return
	}
}

func (h *PostHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	post, err := h.Database.Show(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.Database.Delete(post)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
