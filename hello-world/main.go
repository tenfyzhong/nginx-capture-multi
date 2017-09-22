package main

import (
	"log"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("hello world")
	w.Write([]byte("hello world\n"))
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":8080", nil)
}
