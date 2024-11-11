package cmd

import (
	"github.com/kakky/refacgo/internal/application/evaluation"
	"github.com/kakky/refacgo/internal/presenter"
	loadfile "github.com/kakky/refacgo/pkg/load_file"

	// loadfile "github.com/kakky/refacgo/pkg/load_file"

	"github.com/urfave/cli/v2"
)

type evalCmd struct {
	Evalueation   evaluation.Evaluate
	EvalPresenter presenter.EvalPresenter
}

func newEvalCmd(evaluation evaluation.Evaluate, evalPresenter presenter.EvalPresenter) *evalCmd {
	return &evalCmd{
		Evalueation:   evaluation,
		EvalPresenter: evalPresenter,
	}
}

func (cmd *evalCmd) add() *cli.Command {
	return &cli.Command{
		Name:        "evaluate",
		Aliases:     []string{"eval"},
		Description: "Evaluate code in the specifield file",
		Usage:       "Evaluate code in the specifield file",
		UsageText:   "refacgo eval [option] <filepath>",
		HelpName:    "eval",
		ArgsUsage:   "<filepath> is a path relative to the current directory where the command will be executed",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "japanese",
				Aliases: []string{"j"},
			},
			&cli.StringFlag{
				Name:    "description",
				Aliases: []string{"desc"},
				Value:   "",
				Usage:   "description of code in the specified file",
			},
		},
		Action: cmd.run,
	}
}

func (cmd *evalCmd) run(cCtx *cli.Context) error {
	// 引数のファイルを読み込んで、バイトスライスを格納
	src, err := loadfile.LoadFile(cCtx.Args().Get(0))
	if err != nil {
		return err
	}
	evalb, err := cmd.Evalueation.Evaluate(cCtx.Context, src)
	if err != nil {
		return err
	}
	if err := cmd.EvalPresenter.EvalPrint(cCtx.Context, evalb); err != nil {
		return err
	}
	return nil
}
