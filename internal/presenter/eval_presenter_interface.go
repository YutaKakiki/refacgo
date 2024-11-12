package presenter

import "context"

type EvalPresenter interface {
	EvalPrint(ctx context.Context, eval string) error
}
