package app_test

import (
	"testing"

	"github.com/AidanThomas/ledger/internal/app"
	"github.com/AidanThomas/ledger/internal/domain"
)

func Test_FromConnectionString(t *testing.T) {
	t.Run("postgres - complete", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			User:     "postgres",
			Password: "password",
			Host:     "localhost",
			Database: "database",
			Query:    "sslmode=disable",
			Type:     domain.DBTypePSQL,
			Conn:     "postgres://postgres:password@localhost:5432/database?sslmode=disable",
		}
		actual, err := app.FromConnectionString(expect.Conn)
		if err != nil {
			t.Error(err)
		}
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
	})
}
