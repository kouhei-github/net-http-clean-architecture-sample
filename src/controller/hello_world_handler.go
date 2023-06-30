package controller

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.Header
	fmt.Println(query)
	fmt.Println("TEST")
	fmt.Fprintf(w, "Hello World 1")
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World 2")
}

func Auth(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "auth")
}
