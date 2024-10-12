package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr string
}

var cfg config

func main() {
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// addr := os.Getenv("SNIPPETBOX_ADDR")

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", &home{})

	mux.Handle("GET /home", http.HandlerFunc(homeFunc))

	log.Printf("starting server on %v", cfg.addr)

	err := http.ListenAndServe(cfg.addr, mux)

	log.Fatal(err)
}
