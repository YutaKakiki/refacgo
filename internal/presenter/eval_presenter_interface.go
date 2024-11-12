package presenter

import "context"

type EvalPresenter interface {
	EvalPrint(ctx context.Context, text string) error
}
