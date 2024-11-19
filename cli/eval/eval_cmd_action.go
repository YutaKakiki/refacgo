package eval

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/kakkky/refacgo/cli/shared"
	"github.com/kakkky/refacgo/internal/domain"
	"github.com/kakkky/refacgo/internal/domain/evaluation"
	"github.com/kakkky/refacgo/internal/presenter"
	"github.com/kakkky/refacgo/internal/presenter/indicater"
	"github.com/kakkky/refacgo/pkg/loadfile"

	"github.com/urfave/cli/v2"
)

type evalCmdAction struct {
	evalueation evaluation.Evaluation
	evalPrinter presenter.EvalPrinter
	indicater   indicater.Indicater
}

// cmdActionコンストラクタ
func newEvalCmdAction(evaluation evaluation.Evaluation, evalPrinter presenter.EvalPrinter, indicater indicater.Indicater) *evalCmdAction {
	return &evalCmdAction{
		evalueation: evaluation,
		evalPrinter: evalPrinter,
		indicater:   indicater,
	}
}

// コマンドアクションを初期化する
// japaneseフラグがあれば、日本語対応のEvaluationをDIする
// エントリポイントで初期化したgenAI,evalPrinterもここでDIする
func initEvalCmdAction(cCtx *cli.Context, genAI domain.GenAI, evalPrinter presenter.EvalPrinter, indicater indicater.Indicater) *evalCmdAction {
	var evalCmdAction *evalCmdAction

	if cCtx.Bool("japanese") {
		evalCmdAction = newEvalCmdAction(
			evaluation.NewEvaluationWithGenAiInJap(
				genAI,
			),
			evalPrinter,
			indicater,
		)
	} else {
		evalCmdAction = newEvalCmdAction(
			evaluation.NewEvaluationWithGenAI(
				genAI,
			),
			evalPrinter,
			indicater,
		)
	}
	return evalCmdAction
}

func (eca *evalCmdAction) run(cCtx *cli.Context, ctx context.Context) error {
	eca.indicater.Start()

	if cCtx.NArg() != 1 {
		return errors.New("only one argument, the filename, is required")
	}
	// ファイル名（パス）を引数から取得
	filename := cCtx.Args().Get(0)
	if strings.HasPrefix(filename, `"`) || strings.HasSuffix(filename, `'`) {
		filename = filename[1 : len(filename)-1]
	}
	// 引数のファイルを読み込んで、バイトスライスを格納
	src, err := loadfile.LoadFile(filename)
	if err != nil {
		return err
	}
	// descフラグから文字列を取得し、ソースに追加
	desc := cCtx.String("description")
	// フラグから""が帰ってきた時はそのままソースはそのまま返る
	src = shared.AddDescToSrc(src, desc)
	// Evaluateの結果をモジュール間で逐次出力するためのチャネル
	ch := make(chan string)
	// ビジネスロジック
	// 結果をストリームでチャネルに送信する
	var wg sync.WaitGroup
	wg.Add(1)
	go func() error {
		defer wg.Done()
		err = eca.evalueation.Evaluate(ctx, src, filename, ch)
		if err != nil {
			return err
		}
		return nil
	}()
	<-ch
	eca.indicater.Stop()
	// チャネルからストリームで受信する
	eca.evalPrinter.Print(ctx, ch)
	wg.Wait()
	return nil
}
