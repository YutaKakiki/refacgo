package refactoring

import (
	"context"
	_ "embed"
	"fmt"

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

//go:embed instruction_text/genai_instruction.txt
var instruction string

func (rf *RefactoringWithGenAI) Refactor(ctx context.Context, src []byte, filename string, ch chan<- string) error {
	prompt := fmt.Sprintf("The name of this file is %q.\n\n%v\n\n", filename, instruction)
	if err := rf.genAI.QueryResuluts(ctx, src, prompt, ch); err != nil {
		return err
	}
	return nil
}
