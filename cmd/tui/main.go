package main

import (
	"log"
	"os"

	"github.com/AidanThomas/ledger/config"
	"github.com/AidanThomas/ledger/internal/adapters/tui"
	"github.com/AidanThomas/ledger/internal/app/ledger"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ledger := ledger.New(conf)
	tui := tui.New()

	if err := tui.Run(ledger); err != nil {
		log.Fatal(err)
	}
}
