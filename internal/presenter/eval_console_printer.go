package presenter

import (
	"context"
	"fmt"

	"github.com/kakky/refacgo/internal/presenter/indicater"
)

type EvalConsolePrinter struct {
	indicater *indicater.Indicater
}

func NewEvalConsolePrinter() *EvalConsolePrinter {
	return &EvalConsolePrinter{
		indicater: indicater.NewIndicater(),
	}
}

func (ep *EvalConsolePrinter) Print(ctx context.Context, ch <-chan string) error {
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
