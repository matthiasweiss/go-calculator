package middleware

import (
	"log"
	"net/http"
	"rest-api/common/jwt"
)

func JwtMiddleware(l *log.Logger, s jwt.JwtService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("JWT middleware handles request")

			_, err := s.Verify("Header")

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// TODO: add claims to context
			next.ServeHTTP(w, r)
		})
	}
}
