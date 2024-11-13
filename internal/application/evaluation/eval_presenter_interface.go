package evaluation

import "context"

type EvalPresenter interface {
	EvalPrint(ctx context.Context, ch <-chan string) error
}
