package domain

type Ledger interface {
	Connect(conn string) error
	Execute(query string) (string, error)
}
