package cmd

import (
	"fmt"

	"github.com/kakky/refacgo/internal/application/evaluation"
	"github.com/urfave/cli/v2"
)

type evalCmd struct {
	Evalueation *evaluation.Evaluation
	// EvalPresenter *presenter.EvalPresenter
}

func newEvalCmd(evaluation *evaluation.Evaluation) *evalCmd {
	return &evalCmd{
		Evalueation: evaluation,
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
		Action: func(ctx *cli.Context) error {
			fmt.Println("Evaluate your code !!")
			return nil
		},
	}
}
