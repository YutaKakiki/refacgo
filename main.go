package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kakky/refacgo/cmd"
)

func main() {
	if err := cmd.Execute(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v\n", err)
	}
}
