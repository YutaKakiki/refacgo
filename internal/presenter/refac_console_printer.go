package presenter

import (
	"github.com/kakky/refacgo/internal/presenter/indicater"
)

type RefacConsolePrinter struct {
	inidicater *indicater.Indicater
}

func NewRefacConsolePrinter() *RefacConsolePrinter {
	return &RefacConsolePrinter{
		inidicater: indicater.NewIndicater(),
	}
}

func (ro *RefacConsolePrinter) Print(text string) {

}
