package presenter

import (
	"io"
)

type RefacFileOverWriter struct{}

func NewRefacFileOverWriter() *RefacFileOverWriter {
	return &RefacFileOverWriter{}
}

func (ro *RefacFileOverWriter) OverWrite(w io.Writer, src string) {
}

func (ro *RefacFileOverWriter) OverWriteWithHeaderComment(w io.Writer, src string) {

}
