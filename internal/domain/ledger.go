package domain

type App interface {
	Connect(conn string) error
	Execute(query string) (string, error)
}
