package ports

import "github.com/AidanThomas/ledger/internal/domain"

type ConnectionStore interface {
	Create(conn domain.Connection) error
	ReadAll() ([]domain.Connection, error)
	Update(conn domain.Connection) error
	Delete(conn domain.Connection) error
}
