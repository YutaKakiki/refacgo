package eval

import (
	"context"
	"errors"
	"strings"

	"github.com/kakky/refacgo/cmd/utils"
	"github.com/kakky/refacgo/internal/application"
	"github.com/kakky/refacgo/internal/application/evaluation"
	loadfile "github.com/kakky/refacgo/pkg/load_file"
	"github.com/urfave/cli/v2"
)

type evalCmdAction struct {
	Evalueation   evaluation.Evaluation
	EvalPresenter evaluation.EvalPrinter
}

// cmdActionコンストラクタ
func newEvalCmdAction(evaluation evaluation.Evaluation, evalPresenter evaluation.EvalPrinter) *evalCmdAction {
	return &evalCmdAction{
		Evalueation:   evaluation,
		EvalPresenter: evalPresenter,
	}
}

// コマンドアクションを初期化する
// japaneseフラグがあれば、日本語対応のEvaluationをDIする
// エントリポイントで初期化したgenAI,evalPresenterもここでDIする
func initEvalCmdAction(cCtx *cli.Context, genAI application.GenAI, evalPresenter evaluation.EvalPrinter) *evalCmdAction {
	var evalCmdAction *evalCmdAction

	if cCtx.Bool("japanese") {
		evalCmdAction = newEvalCmdAction(
			evaluation.NewEvaluationWithGenAiInJap(
				genAI,
			),
			evalPresenter,
		)
	} else {
		evalCmdAction = newEvalCmdAction(
			evaluation.NewEvaluationWithGenAI(
				genAI,
			),
			evalPresenter,
		)
	}
	return evalCmdAction
}

func (eca *evalCmdAction) run(cCtx *cli.Context, ctx context.Context) error {
	if cCtx.NArg() != 1 {
		return errors.New("only one argument, the filename, is required")
	}
	// ファイル名（パス）を引数から取得
	filename := cCtx.Args().Get(0)
	if strings.HasPrefix(filename, `"`) || strings.HasSuffix(filename, `'`) {
		filename = filename[1 : len(filename)-1]
	}
	// 引数のファイルを読み込んで、バイトスライスを格納
	src, err := loadfile.LoadExternal(filename)
	if err != nil {
		return err
	}
	// descフラグから文字列を取得し、ソースに追加
	desc := cCtx.String("description")
	// フラグから""が帰ってきた時はそのままソースはそのまま返る
	src = utils.AddDescToSrc(src, desc)
	// Evaluateの結果をモジュール間で逐次出力するためのチャネル
	ch := make(chan string)
	// ビジネスロジック
	// 結果をストリームでチャネルに送信する
	err = eca.Evalueation.Evaluate(ctx, src, filename, ch)
	if err != nil {
		return err
	}
	// チャネルからストリームで受信する
	eca.EvalPresenter.Print(ctx, ch)

	return nil
}
