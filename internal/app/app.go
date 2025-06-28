package app

import (
	"errors"
	"strings"

	"github.com/AidanThomas/ledger/config"
	"github.com/AidanThomas/ledger/internal/adapters/psql"
	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/AidanThomas/ledger/internal/ports"
)

var _ domain.App = (*App)(nil)

type App struct {
	db   ports.Database
	conf config.Configuration
}

func New(conf *config.Configuration) *App {
	return &App{conf: *conf}
}

func (l *App) Connect(conn string) error {
	var dbFlavour string
	for db, prefix := range l.conf.SupportedDBs {
		if strings.HasPrefix(conn, prefix) {
			dbFlavour = db
		}
	}

	var err error
	switch dbFlavour {
	case "psql":
		l.db, err = psql.New(conn)
		if err != nil {
			return err
		}
	default:
		return errors.New("database not supported")
	}
	return nil
}

func (l *App) Execute(query string) (string, error) {
	res, err := l.db.Execute(query)
	if err != nil {
		return "", err
	}

	output := "NO ROWS RETURNED"
	if !res.Empty {
		output = buildTable(*res)
	}
	return output, nil
}

func buildTable(res domain.DBResult) string {
	lengths := make([]int, len(res.Columns))
	modifiedColumns := make([]string, len(res.Columns))
	for i, c := range res.Columns {
		longest := len(c)
		for _, row := range res.Rows {
			l := len(row[i])
			if l > longest {
				longest = l
			}
		}
		lengths[i] = longest
		modifiedColumns[i] = c + strings.Repeat(" ", longest-len(c))
	}
	out := "| " + strings.Join(modifiedColumns, " | ") + " |\n"

	cross := make([]string, len(res.Columns))
	for i, l := range lengths {
		cross[i] = strings.Repeat("-", l)
	}
	out += "|-" + strings.Join(cross, "-|-") + "-|\n"

	var rows []string
	for _, row := range res.Rows {
		modifiedRow := make([]string, len(row))
		for i, v := range row {
			modifiedRow[i] = v + strings.Repeat(" ", lengths[i]-len(v))
		}
		modifiedRow[0] = "| " + modifiedRow[0]
		modifiedRow[len(modifiedRow)-1] += " |"
		rows = append(rows, strings.Join(modifiedRow, " | "))
	}

	out += strings.Join(rows, "\n")
	return out
}
