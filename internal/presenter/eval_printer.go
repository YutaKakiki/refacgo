package presenter

import (
	"context"
	"fmt"
)

type EvalPrinter struct {
	//  プログレスバー
}

func NewEvalPrinter() *EvalPrinter {
	return &EvalPrinter{}
}

func (ep *EvalPrinter) EvalPrint(ctx context.Context, text string) error {
	fmt.Println(text)
	return nil
}
