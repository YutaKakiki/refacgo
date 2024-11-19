package evaluation

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/kakky/refacgo/internal/domain"
)

type EvaluationWithGenAiInJap struct {
	genAI domain.GenAI
}

func NewEvaluationWithGenAiInJap(genAI domain.GenAI) *EvaluationWithGenAiInJap {
	return &EvaluationWithGenAiInJap{
		genAI: genAI,
	}
}

//go:embed instruction_text/genai_instruction_in_jap.txt
var instructionInJap []byte

func (ev *EvaluationWithGenAiInJap) Evaluate(ctx context.Context, src []byte, filename string, ch chan<- string) error {
	prompt := fmt.Sprintf("このファイルの名前は%qです。\n\n%v\n\n", filename, string(instructionInJap))

	go func() error {
		err := ev.genAI.StreamQueryResults(ctx, src, prompt, ch)
		if err != nil {
			return err
		}
		return nil
	}()
	return nil

}