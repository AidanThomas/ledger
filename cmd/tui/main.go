package main

import (
	"log"
	"os"

	"github.com/AidanThomas/ledger/config"
	"github.com/AidanThomas/ledger/internal/app/ledger"
	"github.com/AidanThomas/ledger/internal/app/tui"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ledger := ledger.New(conf)
	ledger.Connect("postgres://postgres:password@localhost:5432/ledger_test?sslmode=disable")
	tui := tui.New(ledger)

	if err := tui.Start(); err != nil {
		log.Fatal(err)
	}
}
