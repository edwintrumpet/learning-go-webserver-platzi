package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello from Go handler")
}

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home works!")
}

func PostRequest(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(writer, "error: %v", err)
		return
	}
	fmt.Fprintf(writer, "Payload %v\n", metadata)
}
