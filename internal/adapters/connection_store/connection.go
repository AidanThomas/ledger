package connection_store

import (
	"errors"
	"fmt"
)

type Connection struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port,omitempty"`
	Database string `json:"database,omitempty"`
	SSLMode  string `json:"sslmode,omitempty"`
	Type     string `json:"type"`
}

func (c *Connection) buildConnectionString() (string, error) {
	switch c.Type {
	case "postgres":
		host := c.Host
		if c.Port != 0 {
			host = fmt.Sprintf("%s:%d", host, c.Port)
		}
		var db string
		if c.Database != "" {
			db = fmt.Sprintf("/%s", c.Database)
		}

		var query string
		if c.SSLMode != "" {
			query += fmt.Sprintf("?sslmode=%s", c.SSLMode)
		}

		return fmt.Sprintf("postgres://%s:%s@%s%s%s", c.User, c.Password, host, db, query), nil
	default:
		return "", errors.New("not supported database type")
	}
}
