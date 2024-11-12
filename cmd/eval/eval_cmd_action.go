package eval

import (
	"github.com/kakky/refacgo/internal/application/evaluation"
	"github.com/kakky/refacgo/internal/config"
	"github.com/kakky/refacgo/internal/gateway/api/gemini"
	"github.com/kakky/refacgo/internal/presenter"
	loadfile "github.com/kakky/refacgo/pkg/load_file"
	"github.com/urfave/cli/v2"
)

type evalCmdAction struct {
	Evalueation   evaluation.Evaluation
	EvalPresenter presenter.EvalPresenter
}

func newEvalCmdAction(evaluation evaluation.Evaluation, evalPresenter presenter.EvalPresenter) *evalCmdAction {
	return &evalCmdAction{
		Evalueation:   evaluation,
		EvalPresenter: evalPresenter,
	}
}

// コマンドアクションを初期化する
// japaneseフラグがあれば、日本語対応のEvaluationをDIする
func initEvalCmdAction(cCtx *cli.Context, cfg *config.Config) *evalCmdAction {
	var evalCmdAction *evalCmdAction

	if cCtx.Bool("japanese") {
		evalCmdAction = newEvalCmdAction(
			evaluation.NewEvaluationWithGenAiInJap(
				gemini.NewGemini(cfg.GeminiConfig, cCtx.Context),
			),
			presenter.NewEvalPrinter(),
		)
	} else {
		evalCmdAction = newEvalCmdAction(
			evaluation.NewEvaluationWithGenAI(
				gemini.NewGemini(cfg.GeminiConfig, cCtx.Context),
			),
			presenter.NewEvalPrinter(),
		)
	}
	return evalCmdAction
}

func (eca *evalCmdAction) run(cCtx *cli.Context) error {
	// ファイル名（パス）を引数から取得
	filename := cCtx.Args().Get(0)
	// 引数のファイルを読み込んで、バイトスライスを格納
	src, err := loadfile.LoadFile(filename)
	if err != nil {
		return err
	}
	// 読み取ったソースファイルを評価する
	eval, err := eca.Evalueation.Evaluate(cCtx.Context, src, filename)
	if err != nil {
		return err
	}
	// 評価コメントをプレゼンターに渡して出力
	if err := eca.EvalPresenter.EvalPrint(cCtx.Context, eval); err != nil {
		return err
	}
	return nil
}
