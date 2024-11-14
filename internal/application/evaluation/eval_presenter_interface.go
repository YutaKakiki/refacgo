package evaluation

import "context"

//go:generate mockgen -package=./evaluation -source=./eval_presenter_interface.go -destination=./eval_presenter_mock.go
type EvalPresenter interface {
	EvalPrint(ctx context.Context, ch <-chan string) error
}
