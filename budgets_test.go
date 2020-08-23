package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/azbshiri/common/test"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var testServer *server

func TestMain(m *testing.M) {
	testServer = newServer(
		pg.Connect(&pg.Options{
			User:     "user",
			Password: "password",
			Database: "demo",
		}),
		mux.NewRouter(),
	)
	os.Exit(m.Run())
}

func TestGetBudgets(t *testing.T) {
	res, err := test.DoRequest(testServer, "GET", "/budgets", nil)
	assert.NoError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)
}
