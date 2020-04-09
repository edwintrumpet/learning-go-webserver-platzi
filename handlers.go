package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello from Go handler")
}

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home works!")
}
