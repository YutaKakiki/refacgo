package refactoring

import (
	"context"

	"github.com/kakky/refacgo/internal/domain"
)

type RefactoringWithGenAI struct {
	genAI domain.GenAI
}

func NewRefactoringWithGenAI(genAI domain.GenAI) *RefactoringWithGenAI {
	return &RefactoringWithGenAI{
		genAI: genAI,
	}
}

func (rf *RefactoringWithGenAI) Refactor(ctx context.Context, src []byte, filename string) (string, error) {
	return "", nil
}
