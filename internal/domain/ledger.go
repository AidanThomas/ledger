package domain

type App interface {
	GetConnections() ([]Connection, error)
	Connect(conn string) error
	Execute(query string) (string, error)
}
