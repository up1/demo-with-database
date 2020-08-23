package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/azbshiri/common/test"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/stretchr/testify/assert"
)

var testServer *server

func TestMain(m *testing.M) {
	testServer = newServer(
		pg.Connect(&pg.Options{
			Addr:     "localhost:5432",
			User:     "user",
			Password: "password",
			Database: "demo",
		}),
		mux.NewRouter(),
	)
	os.Exit(m.Run())
}

func TestGetBudgets_WithSuccessResponse(t *testing.T) {
	var body []budget
	budgets, err := CreateBudgetListFactory(testServer.db, 10)
	assert.NoError(t, err)

	res, err := test.DoRequest(testServer, "GET", BudgetPath, nil)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NoError(t, err)

	ffjson.NewDecoder().DecodeReader(res.Body, &body)
	assert.Len(t, body, 10)
	assert.Equal(t, budgets, &body)
}
