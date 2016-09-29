package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func world(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "World Function!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/world", world)
	http.ListenAndServe(":8001", nil)
}
