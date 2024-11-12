package eval

import (
	"errors"

	"github.com/kakky/refacgo/cmd/utils"
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
	if cCtx.NArg() != 1 {
		return errors.New("Only one argument, the filename, is required")
	}
	// ファイル名（パス）を引数から取得
	filename := cCtx.Args().Get(0)
	// 引数のファイルを読み込んで、バイトスライスを格納
	src, err := loadfile.LoadFile(filename)
	if err != nil {
		return err
	}
	// descフラグから文字列を取得し、ソースに追加
	desc := cCtx.String("description")
	// フラグから""が帰ってきた時はそのままソースはそのまま返る
	src = utils.AddDescToSrc(src, desc)

	// 評価する
	text, err := eca.Evalueation.Evaluate(cCtx.Context, src, filename)
	if err != nil {
		return err
	}
	// 評価コメントをプレゼンターに渡して出力
	if err := eca.EvalPresenter.EvalPrint(cCtx.Context, text); err != nil {
		return err
	}
	return nil
}
