package main

import (
	"embed"
	"github.com/rsturla/golang-aio/internal/cmd"
	"github.com/rsturla/golang-aio/pkg/log"
	"os"
)

//go:embed all:web/dist
var embedFS embed.FS

func main() {
	if err := run(); err != nil {
		log.Errorf("Error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	if err := cmd.Execute(embedFS); err != nil {
		return err
	}

	return nil
}
