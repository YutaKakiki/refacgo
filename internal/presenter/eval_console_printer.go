package presenter

import (
	"context"
	"fmt"
)

type EvalConsolePrinter struct {
}

func NewEvalConsolePrinter() *EvalConsolePrinter {
	return &EvalConsolePrinter{}
}

func (ep *EvalConsolePrinter) Print(ctx context.Context, ch <-chan string) error {
	for text := range ch {
		fmt.Println(text)
	}
	return nil
}
