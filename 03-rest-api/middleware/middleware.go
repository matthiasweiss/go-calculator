package middleware

import (
	"net/http"
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
