package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	log *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	// loggerHandler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(loggerHandler)

	app := &application{
		log: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /download-zip", app.downloadFile)
	mux.HandleFunc("GET /download-image/{path}", app.downloadFileFromUserInput)
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)

	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	logger.Info("starting server", slog.String("addr", ":4000"))

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())

	os.Exit(1)
}
