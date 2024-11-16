package refac

import (
	"context"

	"github.com/kakky/refacgo/internal/application"
	"github.com/kakky/refacgo/internal/application/refactoring"
	"github.com/kakky/refacgo/internal/application/refactoring/diff"
	"github.com/urfave/cli/v2"
)

type refacCmdAction struct {
	refactoring     refactoring.Refactoring
	differ          diff.Differ
	refacPrinter    refactoring.RefacPrinter
	refacOverWriter refactoring.RefacOverWriter
}

func newRefacCmdAction(refactoring refactoring.Refactoring, differ diff.Differ, refacPrinter refactoring.RefacPrinter, refacOverWiter refactoring.RefacOverWriter) *refacCmdAction {
	return &refacCmdAction{
		refactoring:     refactoring,
		differ:          differ,
		refacPrinter:    refacPrinter,
		refacOverWriter: refacOverWiter,
	}
}

func initRefacCmdAction(cCtx *cli.Context, genAI application.GenAI, differ diff.Differ, refacPrinter refactoring.RefacPrinter, refacOverWiter refactoring.RefacOverWriter) *refacCmdAction {
	var refacCmdAction *refacCmdAction

	// -jフラグによってRefacotringインスタンスを切り替える
	// ここでcmdActionを初期化する
	if cCtx.Bool("japanese") {
		// refacCmdAction = newRefacCmdAction(
		// 	refactoring.NewRefactoringWithGenAiInJap(
		// 		genAI,
		// 	),
		// 	diff.NewCmpDiffer(),
		// 	refacPrinter,
		// 	refacOverWiter,
		// )
	} else {
		refacCmdAction = newRefacCmdAction(
			refactoring.NewRefactoringWithGenAI(
				genAI,
			),
			differ,
			refacPrinter,
			refacOverWiter,
		)
	}
	return refacCmdAction
}

func (*refacCmdAction) run(cCtx *cli.Context, ctx context.Context) error {
	return nil
}
