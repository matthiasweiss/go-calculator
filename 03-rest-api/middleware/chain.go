package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Chain struct {
	middlewares []Middleware
}

func NewChain(m ...Middleware) *Chain {
	return &Chain{
		middlewares: m,
	}
}

func (mc *Chain) Apply(handler http.Handler) http.Handler {
	for i := len(mc.middlewares) - 1; i >= 0; i-- {
		handler = mc.middlewares[i](handler)
	}

	return handler
}

func (mc *Chain) Handle(handlerFunc http.HandlerFunc) http.Handler {
	return mc.Apply(handlerFunc)
}
