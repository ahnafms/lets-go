package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/ahnafms/learn-go/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	log      *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Mysql data source name")

	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	// loggerHandler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(loggerHandler)

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
	}

	defer db.Close()

	app := &application{
		log:      logger,
		snippets: &models.SnippetModel{DB: db},
	}

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
