package middleware

import (
	"log"
	"net/http"
)

func JwtMiddleware(next http.Handler, l *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("Jwt verification TODO")
			next.ServeHTTP(w, r)
		})
	}
}
