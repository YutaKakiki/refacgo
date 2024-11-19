package eval

import (
	"context"

	"github.com/kakkky/refacgo/internal/domain"
	"github.com/kakkky/refacgo/internal/presenter"
	"github.com/kakkky/refacgo/internal/presenter/indicater"

	"github.com/urfave/cli/v2"
)

func EvalCmd(ctx context.Context, genAI domain.GenAI, evalPresenter presenter.EvalPrinter, indicater indicater.Indicater) *cli.Command {
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
		Action: func(cCtx *cli.Context) error {
			evalCmdAction := initEvalCmdAction(cCtx, genAI, evalPresenter, indicater)
			if err := evalCmdAction.run(cCtx, ctx); err != nil {
				return err
			}
			return nil
		},
	}
}
