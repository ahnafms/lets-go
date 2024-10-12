package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/ahnafms/learn-go/cmd/internal/api"
	"github.com/ahnafms/learn-go/cmd/internal/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	// loggerHandler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(loggerHandler)

	app := &config.Application{
		Logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", api.Home(app))

	logger.Info("starting server", slog.String("addr", ":4000"))

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())

	os.Exit(1)
}
