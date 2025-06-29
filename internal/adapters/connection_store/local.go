package connection_store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/AidanThomas/ledger/internal/ports"
)

var _ ports.ConnectionStore = (*Local)(nil)

type Local struct {
	fLoc string
}

func NewLocal() (*Local, error) {
	f := os.Getenv("XDG_HOME_DATA")

	if f == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		f = path.Join(homeDir, ".local/share/ledger/ledger_connections.json")
	}

	if err := os.MkdirAll(path.Dir(f), 0700); err != nil {
		return nil, fmt.Errorf("failed to create directory for %s: %w", f, err)
	}

	if _, err := os.Stat(f); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create(f)
		} else {
			return nil, err
		}
	}

	return &Local{
		fLoc: f,
	}, nil
}

func (l *Local) Create(conn domain.Connection) error {
	return nil
}

func (l *Local) ReadAll() ([]domain.Connection, error) {
	file, err := os.ReadFile(l.fLoc)
	if err != nil {
		return nil, fmt.Errorf("failed to read ledger_connections file: %w", err)
	}

	if len(file) == 0 {
		return nil, nil
	}

	var conns []Connection
	if err := json.Unmarshal(file, &conns); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ledger_connections.json: %w", err)
	}

	out := make([]domain.Connection, len(conns))
	for i, c := range conns {
		conn, err := c.buildConnectionString()
		if err != nil {
			return nil, err
		}
		out[i] = domain.Connection{
			ID:   c.ID,
			Name: c.Name,
			Conn: conn,
			Type: c.Type,
		}
	}

	return out, nil
}

func (l *Local) Update(conn domain.Connection) error {
	return nil
}

func (l *Local) Delete(conn domain.Connection) error {
	return nil
}
