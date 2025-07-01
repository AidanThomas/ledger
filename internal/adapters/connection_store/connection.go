package connection_store

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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

func (c *Connection) BuildConnectionString() (string, error) {
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

func (c *Connection) ParseConnectionString(conn string) error {
	switch c.Type {
	case "postgres":
		regex := regexp.MustCompile(`(postgres(?:ql)?):\/\/(?:([^@\s]+)@)?([^\/\s]+)(?:\/(\w+))?(?:\?(.+))?`)
		parts := regex.FindStringSubmatch(conn)
		if parts == nil {
			return errors.New("no matches found")
		}

		login := parts[2]
		server := parts[3]
		db := parts[4]
		query := parts[5]

		c.User = strings.Split(login, ":")[0]
		c.Password = strings.Split(login, ":")[1]
		c.Host = strings.Split(server, ":")[0]
		c.Database = db
		c.SSLMode = query

		return nil
	}
	return nil
}
