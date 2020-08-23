package main

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

type server struct {
	db  *pg.DB
	mux *mux.Router
}

func newServer(db *pg.DB, mux *mux.Router) *server {
	server := server{db, mux}
	server.routes() // register handlers
	return &server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
