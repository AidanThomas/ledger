package app

import (
	"errors"
	"regexp"
	"strings"

	"github.com/AidanThomas/ledger/internal/domain"
)

var (
	psqlRegex = regexp.MustCompile(`(postgres(?:ql)?):\/\/(?:([^@\s]+)@)?([^\/\s]+)(?:\/(\w+))?(?:\?(.+))?`)
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
		out.User = strings.Split(usr, ":")[0]
		out.Password = strings.Split(usr, ":")[1]
		out.Host = strings.Split(srv, ":")[0]
		out.Database = db
		out.Query = qry
		out.Type = domain.DBTypePSQL

		return out, nil
	}

	return domain.Connection{}, errors.New("unsupported db type")
}

func ToConnectionString(c domain.Connection) string {
	return ""
}
