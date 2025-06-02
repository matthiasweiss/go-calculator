package middleware

import (
	"log"
	"net/http"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Jwt verification TODO")
		next.ServeHTTP(w, r)
	})
}
