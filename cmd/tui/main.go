package main

import (
	"log"
	"os"

	"github.com/AidanThomas/ledger/config"
	"github.com/AidanThomas/ledger/internal/adapters/tui"
	"github.com/AidanThomas/ledger/internal/app"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ledger := app.New(conf)
	tui := tui.New(ledger)

	if err := tui.Run(); err != nil {
		log.Fatal(err)
	}
}
