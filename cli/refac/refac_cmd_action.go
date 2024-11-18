package refac

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/kakky/refacgo/cli/refac/utils"
	"github.com/kakky/refacgo/cli/shared"
	"github.com/kakky/refacgo/internal/domain"
	"github.com/kakky/refacgo/internal/domain/refactoring"
	"github.com/kakky/refacgo/internal/domain/refactoring/diff"
	"github.com/kakky/refacgo/internal/presenter"
	"github.com/kakky/refacgo/pkg/loadfile"
	"github.com/urfave/cli/v2"
)

type refacCmdAction struct {
	refactoring     refactoring.Refactoring
	differ          diff.Differ
	refacPrinter    presenter.RefacPrinter
	refacOverWriter presenter.RefacOverWriter
}

func newRefacCmdAction(refactoring refactoring.Refactoring, differ diff.Differ, refacPrinter presenter.RefacPrinter, refacOverWiter presenter.RefacOverWriter) *refacCmdAction {
	return &refacCmdAction{
		refactoring:     refactoring,
		differ:          differ,
		refacPrinter:    refacPrinter,
		refacOverWriter: refacOverWiter,
	}
}

func initRefacCmdAction(cCtx *cli.Context, genAI domain.GenAI, differ diff.Differ, refacPrinter presenter.RefacPrinter, refacOverWiter presenter.RefacOverWriter) *refacCmdAction {
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

func (rca *refacCmdAction) run(cCtx *cli.Context, ctx context.Context) error {
	if cCtx.NArg() != 1 {
		return errors.New("only one argument, the filename, is required")
	}
	// ファイル名（パス）を引数から取得し読み込む
	filename := cCtx.Args().Get(0)
	if strings.HasPrefix(filename, `"`) || strings.HasSuffix(filename, `'`) {
		filename = filename[1 : len(filename)-1]
	}
	originSrc, err := loadfile.LoadFile(filename)
	if err != nil {
		return err
	}
	// descフラグから文字列を取得し、ソースコードに追加
	desc := cCtx.String("description")
	originSrcWithDesc := shared.AddDescToSrc(originSrc, desc)
	// リファクタリングする
	result, err := rca.refactoring.Refactor(ctx, originSrcWithDesc, filename)
	if err != nil {
		return err
	}
	// リファクタリング結果をテキスト/コードに分ける
	code, text, err := utils.DevideCodeAndText(result)
	if err != nil {
		return err
	}
	// 差分を検出
	diff := rca.differ.Diff(string(originSrc), code)
	// ファイル元を書き込み権限で開く
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ファイルへの上書き
	// ヘッダーコメントを付加して上書きする
	rca.refacOverWriter.OverWriteWithHeaderComment(f, code)
	// テキスト・差分を表示
	rca.refacPrinter.Print(text, diff)
	// 上書きを確定するかどうか
	if utils.DecideToApply() {
		rca.refacOverWriter.OverWrite(f, code)
	} else {
		rca.refacOverWriter.OverWrite(f, string(originSrc))
	}
	return nil
}
