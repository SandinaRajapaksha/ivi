package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Handler_1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server\n")
}
func main() {
	//normal funcs converted to handler funcs
	Handler_one := http.HandlerFunc(Handler_1)

	//server mux -  redirect req to correct handler based on pattern
	mux := http.NewServeMux()
	//mux pattern to forward req to handlers ( this comes before making mux timeout)
	mux.HandleFunc("/", Handler_one)
	// timout added
	router := http.TimeoutHandler(mux, 2*time.Second, "request timeout ...")

	//connection - listen from 8080 and to redirect requests to router
	fmt.Println("Listening to port 8080 ...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		errors.New("Cant establish connection")
	}

}
