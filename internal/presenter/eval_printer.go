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

func (ep *EvalPrinter) EvalPrint(ctx context.Context, evalb []byte) error {
	return nil
}
