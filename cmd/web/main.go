package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/AndreReyesG/supermercado/internal/data"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	config   config
	logger   *slog.Logger
	products *data.ProductModel
	//models data.Models
	templateCache map[string]*template.Template
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Web server port")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("SUPERMERCADO_DB_DSN_GO"), "MySQL data source name")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config:   cfg,
		logger:   logger,
		products: &data.ProductModel{DB: db},
		//models: data.NewModels(db),
		templateCache: templateCache,
	}

	logger.Info("starting server", "addr", fmt.Sprintf(":%d", cfg.port))

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
