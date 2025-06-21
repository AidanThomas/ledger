package psql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/AidanThomas/ledger/internal/ports"
	_ "github.com/lib/pq"
)

var _ ports.Database = (*PSQL)(nil)

type PSQL struct {
	ctx context.Context
	db  *sql.DB
}

func New() *PSQL {
	connStr := "postgres://postgres:password@localhost:5432/ledger_test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PSQL{
		ctx: context.Background(),
		db:  db,
	}
}

func (p *PSQL) Execute(query string) (*domain.DBResult, error) {
	rows, err := p.db.QueryContext(p.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// The query doesn't return any rows
	if len(columns) == 0 {
		return &domain.DBResult{Empty: true}, nil
	}

	result := domain.DBResult{Columns: columns}
	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		row := make([]string, len(columns))
		for i, val := range values {
			if val == nil {
				row[i] = "NULL"
			} else {
				row[i] = fmt.Sprintf("%v", val)
			}
		}

		result.Rows = append(result.Rows, row)
	}

	return &result, nil
}
