package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (router *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exists := router.rules[path]
	handler, methodExists := router.rules[path][method]
	return handler, methodExists, exists
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, methodExists, exists := router.FindHandler(request.URL.Path, request.Method)
	if !exists {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExists {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(writer, request)
}
