package router

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct{}

type HttpMethod int

const (
	Get HttpMethod = iota
	Post
	Put
	Patch
	Delete
)

var httpMethodNames = map[HttpMethod]string{
	Get:    "GET",
	Post:   "POST",
	Put:    "PUT",
	Patch:  "PATCH",
	Delete: "DELETE",
}

func (method HttpMethod) String() string {
	if name, ok := httpMethodNames[method]; ok {
		return name
	}

	return httpMethodNames[HttpMethod(method)]
}

func New() *Router {
	return &Router{}
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle(Get, path, handler)
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle(Post, path, handler)
}

func (r *Router) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle(Put, path, handler)
}

func (r *Router) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle(Patch, path, handler)
}

func (r *Router) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle(Delete, path, handler)
}

func (r *Router) handle(method HttpMethod, path string, handler func(http.ResponseWriter, *http.Request)) {
	prefix, err := prefix(method)

	if err != nil {
		log.Fatalf("Invalid method: %s", method)
	}

	http.HandleFunc(fmt.Sprintf("%s %s", prefix, path), handler)
}

func prefix(method HttpMethod) (string, error) {
	if val, ok := httpMethodNames[method]; ok {
		return val, nil
	}

	return "", fmt.Errorf("unmapped HTTP method: %d", method)
}
