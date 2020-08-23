package main

import (
	"math/rand"
	"time"

	"github.com/go-pg/pg"
)

func CreateBudgetListFactory(db *pg.DB, length int) (*[]budget, error) {
	budgets := make([]budget, length)
	for _, budget := range budgets {
		budget.Amount = rand.Float64()
		budget.CreatedAt = time.Now()
	}

	err := db.Insert(&budgets)
	if err != nil {
		return nil, err
	}

	return &budgets, nil
}
