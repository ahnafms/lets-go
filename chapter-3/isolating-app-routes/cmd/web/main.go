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

	logger.Info("starting server", slog.String("addr", ":4000"))

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
