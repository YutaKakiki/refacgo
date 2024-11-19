package main

import (
	"context"
	"log"

	"github.com/kakky/refacgo/cli"
	"github.com/kakky/refacgo/internal/config"
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
