package refactoring

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/kakky/refacgo/internal/domain"
)

type RefactoringWithGenAiInJap struct {
	genAI domain.GenAI
}

func NewRefactoringWithGenAiInJap(genAI domain.GenAI) *RefactoringWithGenAiInJap {
	return &RefactoringWithGenAiInJap{
		genAI: genAI,
	}
}

//go:embed instruction_text/genai_instruction_in_jap.txt
var instructionInJap string

func (rf *RefactoringWithGenAiInJap) Refactor(ctx context.Context, src []byte, filename string, ch chan<- string) error {
	prompt := fmt.Sprintf("このファイル名は %q です。\n\n%v\n\n", filename, instructionInJap)
	if err := rf.genAI.QueryResuluts(ctx, src, prompt, ch); err != nil {
		return err
	}
	return nil
}
