package presenter

import (
	"io"

	"github.com/kakky/refacgo/internal/presenter/indicater"
)

type RefacFileOverWriter struct {
	inidicater *indicater.Indicater
}

func NewRefacFileOverWriter() *RefacFileOverWriter {
	return &RefacFileOverWriter{
		inidicater: indicater.NewIndicater(),
	}
}

func (ro *RefacFileOverWriter) OverWrite(w io.Writer, src string) {
}
