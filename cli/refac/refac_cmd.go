package refac

import (
	"context"

	"github.com/kakky/refacgo/internal/domain"
	"github.com/kakky/refacgo/internal/domain/refactoring/diff"
	"github.com/kakky/refacgo/internal/presenter"
	"github.com/kakky/refacgo/internal/presenter/indicater"
	"github.com/urfave/cli/v2"
)

func RefacCmd(ctx context.Context, genAI domain.GenAI, differ diff.Differ, refacPrinter presenter.RefacPrinter,
	refacOverWiter presenter.RefacOverWriter, indicater indicater.Indicater) *cli.Command {
	return &cli.Command{
		Name:        "refactor",
		Aliases:     []string{"refac"},
		Description: "refactor code in the specifield file",
		Usage:       "refactor code in the specifield file",
		UsageText:   "refacgo refac [option] <filepath>",
		HelpName:    "refac",
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
			refacCmdAction := initRefacCmdAction(cCtx, genAI, differ, refacPrinter, refacOverWiter, indicater)
			if err := refacCmdAction.run(cCtx, ctx); err != nil {
				return err
			}
			return nil
		},
	}
}
