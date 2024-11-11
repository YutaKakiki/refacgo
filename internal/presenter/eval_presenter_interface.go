package presenter

import "context"

type EvalPresenter interface {
	EvalPrint(ctx context.Context, evalb []byte) error
}
