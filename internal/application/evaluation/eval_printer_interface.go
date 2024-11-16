package evaluation

import (
	"context"
)

//go:generate mockgen -package=evaluation -source=./eval_printer_interface.go -destination=./eval_printer_mock.go
type EvalPrinter interface {
	Print(ctx context.Context, ch <-chan string) error
}
