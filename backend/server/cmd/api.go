package main

import (
	"errors"
	"net/http"
)

func main() {
	//server mux - router that redirect req to correct handler based on pattern
	router := http.NewServeMux()

	//connection - listen from 8080 and to redirect requests to router
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		errors.New("Cant establish connection")
	}
}
