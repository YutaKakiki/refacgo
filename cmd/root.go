package cmd

import (
	"context"
	"os"

	"github.com/kakky/refacgo/cmd/eval"
	"github.com/kakky/refacgo/internal/config"
	"github.com/urfave/cli/v2"
)

const (
	version = "v1.0"
)

func Execute(ctx context.Context, cfg *config.Config) error {
	app := &cli.App{
		Name:        "refacgo",
		Version:     version,
		Description: "A Go-based command-line tool that evaluates the code in a specified Go file and provides refactoring suggestions powered by AI",
		Commands: []*cli.Command{
			eval.EvalCmd(cfg),
		},
	}
	if err := app.RunContext(ctx, os.Args); err != nil {
		return err
	}
	return nil
}
