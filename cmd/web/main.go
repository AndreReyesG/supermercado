package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server on :4000")

	err := http.ListenAndServe(":4000", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
