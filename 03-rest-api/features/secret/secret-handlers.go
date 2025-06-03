package secret

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SecretHandlers struct {
	Logger *log.Logger
}

func NewSecretHandlers(l *log.Logger) *SecretHandlers {
	return &SecretHandlers{
		Logger: l,
	}
}

func (h *SecretHandlers) Index(w http.ResponseWriter, r *http.Request) {
	secret := secret{
		Message: "Hello there",
	}

	err := json.NewEncoder(w).Encode(secret)

	if err != nil {
		http.Error(w, "Could not encode secret", http.StatusBadRequest)
		return
	}
}

func (h *SecretHandlers) Show(w http.ResponseWriter, r *http.Request) {
	wildcard := r.PathValue("wildcard")

	secret := secret{
		Message: fmt.Sprintf("Hello secret %s", wildcard),
	}

	err := json.NewEncoder(w).Encode(secret)

	if err != nil {
		http.Error(w, "Could not encode secret", http.StatusBadRequest)
		return
	}
}
