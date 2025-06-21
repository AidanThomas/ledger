package ports

import "github.com/AidanThomas/ledger/internal/domain"

type Database interface {
	Execute(query string) (*domain.DBResult, error)
}
