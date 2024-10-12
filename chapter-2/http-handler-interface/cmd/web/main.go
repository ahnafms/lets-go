package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", &home{})

	mux.Handle("GET /home", http.HandlerFunc(homeFunc))

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
