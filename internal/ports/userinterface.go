package ports

import "github.com/AidanThomas/ledger/internal/domain"

type UserInterface interface {
	Run(domain.Ledger) error
}
