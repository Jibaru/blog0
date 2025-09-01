package handler

import (
	"net/http"

	"blog0/config"
	"blog0/db"
	"blog0/server"
)

var cfg = config.Load()

func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := db.New(cfg.PostgresURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	server.New(cfg, db).ServeHTTP(w, r)
}
