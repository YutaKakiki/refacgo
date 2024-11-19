package refac

import (
	"context"
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/kakky/refacgo/cli/refac/utils"
	"github.com/kakky/refacgo/cli/shared"
	"github.com/kakky/refacgo/internal/domain"
	"github.com/kakky/refacgo/internal/domain/refactoring"
	"github.com/kakky/refacgo/internal/domain/refactoring/diff"
	"github.com/kakky/refacgo/internal/presenter"
	"github.com/kakky/refacgo/internal/presenter/indicater"
	"github.com/kakky/refacgo/pkg/loadfile"
	"github.com/urfave/cli/v2"
)

type refacCmdAction struct {
	refactoring     refactoring.Refactoring
	differ          diff.Differ
	refacPrinter    presenter.RefacPrinter
	refacOverWriter presenter.RefacOverWriter
	indicater       indicater.Indicater
}

func newRefacCmdAction(refactoring refactoring.Refactoring, differ diff.Differ, refacPrinter presenter.RefacPrinter,
	refacOverWiter presenter.RefacOverWriter, indicater indicater.Indicater) *refacCmdAction {
	return &refacCmdAction{
		refactoring:     refactoring,
		differ:          differ,
		refacPrinter:    refacPrinter,
		refacOverWriter: refacOverWiter,
		indicater:       indicater,
	}
}

func initRefacCmdAction(cCtx *cli.Context, genAI domain.GenAI, differ diff.Differ, refacPrinter presenter.RefacPrinter,
	refacOverWiter presenter.RefacOverWriter, indicater indicater.Indicater) *refacCmdAction {
	var refacCmdAction *refacCmdAction

	// -jフラグによってRefacotringインスタンスを切り替える
	// ここでcmdActionを初期化する
	if cCtx.Bool("japanese") {
		refacCmdAction = newRefacCmdAction(
			refactoring.NewRefactoringWithGenAiInJap(
				genAI,
			),
			diff.NewCmpDiffer(),
			refacPrinter,
			refacOverWiter,
			indicater,
		)
	} else {
		refacCmdAction = newRefacCmdAction(
			refactoring.NewRefactoringWithGenAI(
				genAI,
			),
			differ,
			refacPrinter,
			refacOverWiter,
			indicater,
		)
	}
	return refacCmdAction
}

func (rca *refacCmdAction) run(cCtx *cli.Context, ctx context.Context) error {
	// インジケータースピナを回す
	rca.indicater.Start()
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
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() error {
		defer wg.Done()
		if err := rca.refactoring.Refactor(ctx, originSrcWithDesc, filename, ch); err != nil {
			return err
		}
		return nil
	}()
	var result string
	for s := range ch {
		result += s
	}
	// リファクタリング結果をテキスト/コードに分ける
	code, text, err := utils.DevideCodeAndText(result)
	if err != nil {
		return err
	}
	// 差分を検出
	diff := rca.differ.Diff(string(originSrc), code)
	// インジケータースピナーを止める
	rca.indicater.Stop()
	// テキスト・差分を表示
	rca.refacPrinter.Print(text, diff)
	// ファイルへの上書き
	// ヘッダーコメントを付加して上書きする
	rca.refacOverWriter.OverWriteWithHeaderComment(filename, code)
	// 上書きを確定するかどうか
	if utils.DecideToApply(os.Stdin) {
		rca.refacOverWriter.OverWrite(filename, code)
	} else {
		rca.refacOverWriter.OverWrite(filename, string(originSrc))
	}
	wg.Wait()
	return nil
}
