package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/home", http.HandlerFunc(hello))
	http.ListenAndServe(":4000", nil)
}
