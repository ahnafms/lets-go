package main

import (
	"net/http"
)

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
