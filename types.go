package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Metadata interface{}

type User struct {
	name  string
	email string
	phone string
}
