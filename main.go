package main

import (
	"context"
	"log"

	"github.com/kakkky/refacgo/cli"
	"github.com/kakkky/refacgo/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	if err := cli.Execute(context.Background(), cfg); err != nil {
		log.Fatalf("Error occured in executing command: %v", err)
	}
}
