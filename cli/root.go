package cli

import (
	"context"
	"os"

	"github.com/kakky/refacgo/cli/eval"
	"github.com/kakky/refacgo/cli/refac"
	"github.com/kakky/refacgo/internal/config"
	"github.com/kakky/refacgo/internal/domain/refactoring/diff"
	"github.com/kakky/refacgo/internal/gateway/api/gemini"
	"github.com/kakky/refacgo/internal/presenter"
	"github.com/kakky/refacgo/internal/presenter/indicater"
	"github.com/urfave/cli/v2"
)

const (
	version = "v1.0"
)

func Execute(ctx context.Context, cfg *config.Config) error {
	// geminiを初期化
	gemini := gemini.NewGemini(cfg.GeminiConfig, ctx)
	app := &cli.App{
		Name:        "refacgo",
		Version:     version,
		Description: "A Go-based command-line tool that evaluates the code in a specified Go file and provides refactoring suggestions powered by AI",
		Commands: []*cli.Command{
			eval.EvalCmd(ctx, gemini, presenter.NewEvalConsolePrinter(), indicater.NewEvalSpinner()),
			refac.RefacCmd(ctx, gemini, diff.NewCmpDiffer(), presenter.NewRefacConsolePrinter(), presenter.NewRefacFileOverWriter(), indicater.NewRefacSpinner()),
		},
	}
	if err := app.RunContext(ctx, os.Args); err != nil {
		return err
	}
	return nil
}
