package main

import (
	"encoding/json"
	"net/http"

	"github.com/azbshiri/common/db"
	"github.com/pquerna/ffjson/ffjson"
)

type budget struct {
	db.Model
	Amount float64 `json:"amount"`
}

func (s *server) getBudgets(w http.ResponseWriter, r *http.Request) {
	var budgets []*budget
	err := s.db.Model(&budgets).Select()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(DatabaseError)
		return
	}

	ffjson.NewEncoder(w).Encode(budgets)
}
