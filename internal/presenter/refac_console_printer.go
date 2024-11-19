package presenter

import (
	"fmt"
)

type RefacConsolePrinter struct {
}

func NewRefacConsolePrinter() *RefacConsolePrinter {
	return &RefacConsolePrinter{}
}

func (ro *RefacConsolePrinter) Print(text ...string) {
	for _, t := range text {
		fmt.Println(t)
	}
}
