package presenter

import (
	"context"
)

type EvalPrinter struct {
	//  プログレスバー
}

func NewEvalPrinter() *EvalPrinter {
	return &EvalPrinter{}
}

func (ep *EvalPrinter) EvalPrint(ctx context.Context, eval string) error {
	return nil
}
