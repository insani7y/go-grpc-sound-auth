package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/reqww/go-rest-api/internal/app/store/sql_store"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	defer db.Close()
	if err != nil {
		return err
	}
	store := sql_store.New(db)
	server := newServer(store)
	server.logger.Info(fmt.Sprintf("HTTP server started at %v", config.BindAddr))
	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}