package eval

import (
	"context"

	"github.com/kakky/refacgo/internal/config"

	"github.com/urfave/cli/v2"
)

func EvalCmd(cfg *config.Config, ctx context.Context) *cli.Command {
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
			evalCmdAction := initEvalCmdAction(cCtx, cfg)
			if err := evalCmdAction.run(cCtx, ctx); err != nil {
				return err
			}
			return nil
		},
	}
}
