package domain

type App interface {
	GetConnections() ([]Connection, error)
	AddConnection(c Connection) error
	Connect(conn string) error
	Execute(query string) (string, error)
}
