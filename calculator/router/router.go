package router

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

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
	return &Router{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
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
	methodPrefix, err := prefix(method)

	if err != nil {
		log.Fatalf("Invalid method: %s", method)
		return
	}

	key := fmt.Sprintf("%s %s", methodPrefix, path)
	r.routes[key] = handler
}

func prefix(method HttpMethod) (string, error) {
	if val, ok := httpMethodNames[method]; ok {
		return val, nil
	}

	return "", fmt.Errorf("unmapped HTTP method: %d", method)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s %s", req.Method, req.URL.Path)

	if handler, ok := r.routes[key]; ok {
		handler(w, req)
		return
	}

	http.NotFound(w, req)
}
