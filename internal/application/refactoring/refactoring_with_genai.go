package refactoring

import (
	"context"

	"github.com/kakky/refacgo/internal/application"
)

type RefactoringWithGenAI struct {
	genAI application.GenAI
}

func NewRefactoringWithGenAI(genAI application.GenAI) *RefactoringWithGenAI {
	return &RefactoringWithGenAI{
		genAI: genAI,
	}
}

func (rf *RefactoringWithGenAI) Refactor(ctx context.Context, src []byte, filename string, ch chan<- string) ([]byte, error) {
	return nil, nil
}
