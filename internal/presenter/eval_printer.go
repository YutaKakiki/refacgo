package presenter

import (
	"context"
	"fmt"

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
	<-ch            // チャネルから受信するまではブロッキング
	is.Stop()
	for text := range ch {
		fmt.Println(text)
	}
	return nil
}
