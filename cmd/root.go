package cmd

import (
	"context"
	"os"

	"github.com/kakky/refacgo/internal/application/evaluation"
	"github.com/kakky/refacgo/internal/presenter"
	"github.com/urfave/cli/v2"
)

const (
	version = "v1.0"
)

func Execute(ctx context.Context) error {
	evalCmd := newEvalCmd(evaluation.NewEvaluateByGemini(), presenter.NewEvalPrinter())
	// refactorCmd:=newRefactorCmd(refactoring.NewRefactoring())
	app := &cli.App{
		Name:        "refacgo",
		Version:     version,
		Description: "A Go-based command-line tool that evaluates the code in a specified Go file and provides refactoring suggestions powered by AI",
		Commands: []*cli.Command{
			evalCmd.add(),
			// refactorCmd.add()
		},
	}
	if err := app.RunContext(ctx, os.Args); err != nil {
		return err
	}
	return nil
}
