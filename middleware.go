package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(nextFunction http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := true
			fmt.Println("Checking auth")
			if flag {
				nextFunction(writer, request)
			} else {
				return
			}
		}
	}
}
