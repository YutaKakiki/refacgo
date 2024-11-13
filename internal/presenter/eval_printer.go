package presenter

import (
	"context"
	"fmt"
	"time"

	"github.com/kakky/refacgo/internal/presenter/indicater"
)

type EvalPrinter struct {
	indicater *indicater.Indicater
}

func NewEvalPrinter() *EvalPrinter {
	return &EvalPrinter{
		indicater: indicater.NewIndicater(),
	}
}

func (ep *EvalPrinter) EvalPrint(ctx context.Context, ch <-chan string) error {
	is := ep.indicater.Spinner
	is.Suffix = "  Waiting for evaluating..."
	is.Start()
	defer is.Stop() // 処理の最後に必ずスピナーを停止する

	for {
		select {
		case <-ctx.Done():
			return ctx.Err() // キャンセル通知がされた場合
		case text, ok := <-ch:
			if !ok {
				return nil // チャネルが閉じられた場合
			}
			if is.Active() {
				time.Sleep(2 * time.Second)
				is.Stop()
			}
			fmt.Println(text)
		}
	}
}
