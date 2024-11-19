package presenter

import (
	_ "embed"
	"fmt"
	"os"
)

type RefacFileOverWriter struct{}

func NewRefacFileOverWriter() *RefacFileOverWriter {
	return &RefacFileOverWriter{}
}

func (ro *RefacFileOverWriter) OverWrite(filename, src string) error {
	// ファイル元を書き込み権限で開く
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(src))
	if err != nil {
		return err
	}
	return nil
}

//go:embed comment/header_comment.txt
var headerComment string

func (ro *RefacFileOverWriter) OverWriteWithHeaderComment(filename, src string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	srcWithComment := fmt.Sprintf("%v\n\n\n%v", headerComment, src)
	_, err = f.Write([]byte(srcWithComment))
	if err != nil {
		return err
	}
	return nil
}
