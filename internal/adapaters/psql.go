package psql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PSQL struct {
	ctx context.Context
	db  *sql.DB
}

func NewPSQL() *PSQL {
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

func (p *PSQL) DoQuery(query string) ([]string, [][]string, error) {
	rows, err := p.db.QueryContext(p.ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	// The query doesn't return any rows
	if len(columns) == 0 {
		return nil, nil, nil
	}

	var result [][]string
	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, nil, err
		}

		row := make([]string, len(columns))
		for i, val := range values {
			if val == nil {
				row[i] = "NULL"
			} else {
				row[i] = fmt.Sprintf("%v", val)
			}
		}

		result = append(result, row)

	}

	return columns, result, nil
}
