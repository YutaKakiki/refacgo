package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kakky/refacgo/cmd"
	"github.com/kakky/refacgo/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse config: %v", err)
	}
	if err := cmd.Execute(context.Background(), cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error occured in executing command : %v\n", err)
	}
}
