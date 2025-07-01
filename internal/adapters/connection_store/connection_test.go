package connection_store_test

import (
	"fmt"
	"testing"

	"github.com/AidanThomas/ledger/internal/adapters/connection_store"
)

func TestConnection(t *testing.T) {
	t.Run("parse postgres connection", func(t *testing.T) {
		c := &connection_store.Connection{
			ID:   1,
			Name: "test",
			Type: "postgres",
		}
		if err := c.ParseConnectionString("postgres://postres:password@localhost:5432/database?sslmode=disable"); err == nil {
			t.Error(err)
		}

		fmt.Println(c)
	})
}
