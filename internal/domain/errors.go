package domain

import "errors"

var (
	ErrUnsupportedDb = errors.New("unsupported database type")
)
