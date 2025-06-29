package config

import (
	"os"
)

type Configuration struct {
	SupportedDBs map[string]string
	StateDir     string
}

func Load() (*Configuration, error) {
	var conf Configuration

	conf.SupportedDBs = map[string]string{
		"psql": "postgres",
	}

	lState := os.Getenv("XDG_DATA_HOME")
	if lState == "" {
		conf.StateDir = "~/.local/share/ledger"
	} else {
		conf.StateDir = lState
	}

	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Configuration) Validate() error {
	return nil
}
