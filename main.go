package main

import (
	"embed"
	"os"

	"github.com/rsturla/golang-aio/internal/cmd"
	"github.com/rsturla/golang-aio/pkg/log"
)

//go:embed all:web/dist
var embedFS embed.FS

func main() {
	if err := cmd.Execute(embedFS); err != nil {
		log.Fatalf("Run failed with error: %s", err)
		os.Exit(1)
	}
}
