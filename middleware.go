package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

func Logging() Middleware {
	return func(nextFunction http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			jobStart := time.Now()
			defer func() {
				log.Println(request.URL.Path, time.Since(jobStart))
			}()
			nextFunction(writer, request)
		}
	}
}
