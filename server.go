package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (server *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := server.router.rules[path]
	if !exists {
		server.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	server.router.rules[path][method] = handler
}

func (server *Server) AddMiddleware(nextFunction http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, element := range middlewares {
		nextFunction = element(nextFunction)
	}
	return nextFunction
}

func (server *Server) Listen() error {
	http.Handle("/", server.router)
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}
	return nil
}
