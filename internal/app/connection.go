package app

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AidanThomas/ledger/internal/domain"
)

var (
	psqlRegex = regexp.MustCompile(`(postgres(?:ql)?):\/\/(?:([^@\s]+)@)?([^\/\?\s]+)(?:\/(\w+))?(?:\?(.+))?`)
)

func FromConnectionString(conn string) (domain.Connection, error) {
	if match := psqlRegex.FindStringSubmatch(conn); match != nil {
		out := domain.Connection{Conn: conn}
		sch := match[1]
		usr := match[2]
		srv := match[3]
		db := match[4]
		qry := match[5]

		out.Scheme = sch
		usrParts := strings.Split(usr, ":")
		out.User = usrParts[0]
		if len(usrParts) > 1 {
			out.Password = usrParts[1]
		}
		out.Host = srv
		out.Database = db
		out.Query = qry
		out.Type = domain.DBTypePSQL

		return out, nil
	}

	return domain.Connection{}, domain.ErrUnsupportedDb
}

func ToConnectionString(c domain.Connection) (string, error) {
	if c.Conn != "" {
		return c.Conn, nil
	}

	switch c.Type {
	case domain.DBTypePSQL:
		return buildPSQLConnectionString(c), nil
	}

	return "", domain.ErrUnsupportedDb
}

func buildPSQLConnectionString(c domain.Connection) string {
	out := fmt.Sprintf("%s://", c.Scheme)
	login := ""
	if c.User != "" {
		login += fmt.Sprintf("%s", c.User)
	}
	if c.Password != "" {
		login += fmt.Sprintf(":%s", c.Password)
	}
	if login != "" {
		out += fmt.Sprintf("%s@", login)
	}
	out += c.Host
	if c.Database != "" {
		out += fmt.Sprintf("/%s", c.Database)
	}
	if c.Query != "" {
		out += fmt.Sprintf("?%s", c.Query)
	}
	return out
}
