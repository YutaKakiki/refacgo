package evaluation

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/kakky/refacgo/internal/domain"
)

type EvaluationWithGenAI struct {
	genAI domain.GenAI
}

func NewEvaluationWithGenAI(genAI domain.GenAI) *EvaluationWithGenAI {
	return &EvaluationWithGenAI{
		genAI: genAI,
	}
}

//go:embed instruction_text/genai_instruction.txt
var instruction []byte

func (ev *EvaluationWithGenAI) Evaluate(ctx context.Context, src []byte, filename string, ch chan<- string) error {
	prompt := fmt.Sprintf("The name of this file is %q.\n\n%v\n\n", filename, string(instruction))
	go func() error {
		err := ev.genAI.StreamQueryResults(ctx, src, prompt, ch)
		if err != nil {
			return err
		}
		return nil
	}()
	return nil
}
