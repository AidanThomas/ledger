package app_test

import (
	"testing"

	"github.com/AidanThomas/ledger/internal/app"
	"github.com/AidanThomas/ledger/internal/domain"
)

func Test_FromConnectionString(t *testing.T) {
	t.Run("postgres - all", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			User:     "postgres",
			Password: "password",
			Host:     "localhost:5432",
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
			t.Errorf("\nExpected: %v,\nActual:   %v", expect, actual)
		}
	})
	t.Run("postgres - no login", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			Host:     "localhost:5432",
			Database: "database",
			Query:    "sslmode=disable",
			Type:     domain.DBTypePSQL,
			Conn:     "postgres://localhost:5432/database?sslmode=disable",
		}
		actual, err := app.FromConnectionString(expect.Conn)
		if err != nil {
			t.Error(err)
		}
		if actual != expect {
			t.Errorf("\nExpected: %v,\nActual:   %v", expect, actual)
		}
	})
	t.Run("postgres - no password", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			User:     "postgres",
			Host:     "localhost:5432",
			Database: "database",
			Query:    "sslmode=disable",
			Type:     domain.DBTypePSQL,
			Conn:     "postgres://postgres@localhost:5432/database?sslmode=disable",
		}
		actual, err := app.FromConnectionString(expect.Conn)
		if err != nil {
			t.Error(err)
		}
		if actual != expect {
			t.Errorf("\nExpected: %v,\nActual:   %v", expect, actual)
		}
	})
	t.Run("postgres - no database", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			User:     "postgres",
			Password: "password",
			Host:     "localhost:5432",
			Query:    "sslmode=disable",
			Type:     domain.DBTypePSQL,
			Conn:     "postgres://postgres:password@localhost:5432?sslmode=disable",
		}
		actual, err := app.FromConnectionString(expect.Conn)
		if err != nil {
			t.Error(err)
		}
		if actual != expect {
			t.Errorf("\nExpected: %v,\nActual:   %v", expect, actual)
		}
	})
	t.Run("postgres - no query", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			User:     "postgres",
			Password: "password",
			Host:     "localhost:5432",
			Database: "database",
			Type:     domain.DBTypePSQL,
			Conn:     "postgres://postgres:password@localhost:5432/database",
		}
		actual, err := app.FromConnectionString(expect.Conn)
		if err != nil {
			t.Error(err)
		}
		if actual != expect {
			t.Errorf("\nExpected: %v,\nActual:   %v", expect, actual)
		}
	})
	t.Run("postgres - no port", func(t *testing.T) {
		expect := domain.Connection{
			Scheme:   "postgres",
			User:     "postgres",
			Password: "password",
			Host:     "localhost",
			Database: "database",
			Type:     domain.DBTypePSQL,
			Conn:     "postgres://postgres:password@localhost/database",
		}
		actual, err := app.FromConnectionString(expect.Conn)
		if err != nil {
			t.Error(err)
		}
		if actual != expect {
			t.Errorf("\nExpected: %v,\nActual:   %v", expect, actual)
		}
	})
}
