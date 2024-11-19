package cli

import (
	"context"
	"os"

	"github.com/kakkky/refacgo/cli/eval"
	"github.com/kakkky/refacgo/cli/refac"
	"github.com/kakkky/refacgo/internal/config"
	"github.com/kakkky/refacgo/internal/domain/refactoring/diff"
	"github.com/kakkky/refacgo/internal/gateway/api/gemini"
	"github.com/kakkky/refacgo/internal/presenter"
	"github.com/kakkky/refacgo/internal/presenter/indicater"
	"github.com/urfave/cli/v2"
)

const (
	version = "v0.1.0"
)

func Execute(ctx context.Context, cfg *config.Config) error {
	// geminiを初期化
	gemini := gemini.NewGemini(cfg.GeminiConfig, ctx)
	app := &cli.App{
		Name:                 "refacgo",
		Version:              version,
		Description:          "A Go-based command-line tool that evaluates the code in a specified Go file and provides refactoring suggestions powered by AI",
		EnableBashCompletion: true,
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
