package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareChain struct {
	middlewares []Middleware
}

func NewMiddlewareChain(m ...Middleware) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: m,
	}
}

func (mc *MiddlewareChain) Apply(handler http.Handler) http.Handler {
	for i := len(mc.middlewares) - 1; i >= 0; i-- {
		handler = mc.middlewares[i](handler)
	}

	return handler
}

func (mc *MiddlewareChain) Handle(handlerFunc http.HandlerFunc) http.Handler {
	return mc.Apply(handlerFunc)
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		duration := time.Since(start)
		log.Printf("Timestamp: %s, Status: %d, Method: %s, Path: %s, Duration: %v",
			start.Format(time.RFC3339),
			recorder.statusCode,
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}

func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
