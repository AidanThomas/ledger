package domain

type DBRow []string

type DBResult struct {
	Columns []string
	Rows    []DBRow
	Empty   bool
}
