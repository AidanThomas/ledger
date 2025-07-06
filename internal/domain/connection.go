package domain

type DBType string

const (
	DBTypePSQL = "psql"
)

type Connection struct {
	ID       int    `json:"id"`
	Scheme   string `json:"scheme"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database,omitempty"`
	Query    string `json:"query,omitempty"`
	Type     DBType `json:"type"`
	Conn     string `json:"-"`
}
