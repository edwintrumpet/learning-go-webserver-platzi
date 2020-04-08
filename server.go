package main

import "net/http"

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (server *Server) Listen() error {
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}
	return nil
}
