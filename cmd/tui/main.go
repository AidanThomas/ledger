package main

import (
	"log"
	"os"

	"github.com/AidanThomas/ledger/config"
	"github.com/AidanThomas/ledger/internal/adapters/connection_store"
	"github.com/AidanThomas/ledger/internal/adapters/ui/tui"
	"github.com/AidanThomas/ledger/internal/app"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cs, err := connection_store.NewLocal()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ledger := app.New(conf, cs)
	tui := tui.New(ledger)

	if err := tui.Run(); err != nil {
		log.Fatal(err)
	}
}
