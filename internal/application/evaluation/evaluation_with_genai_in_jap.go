package evaluation

import (
	"context"
	"fmt"

	"github.com/kakky/refacgo/internal/application"
	loadfile "github.com/kakky/refacgo/pkg/load_file"
)

type EvaluationWithGenAiInJap struct {
	genAI application.GenAI
}

func NewEvaluationWithGenAiInJap(genAI application.GenAI) *EvaluationWithGenAiInJap {
	return &EvaluationWithGenAiInJap{
		genAI: genAI,
	}
}

func (ev *EvaluationWithGenAiInJap) Evaluate(ctx context.Context, src []byte, filename string, ch chan<- string) error {
	instruction, err := loadfile.LoadInternal("./instruction_text/genai_instruction_in_jap.txt")
	if err != nil {
		panic(err)
	}
	prompt := fmt.Sprintf("このファイルの名前は%qです。\n\n%v\n\n", filename, string(instruction))

	go func() error {
		err = ev.genAI.StreamQueryResults(ctx, src, prompt, ch)
		if err != nil {
			return err
		}
		return nil
	}()
	return nil

}
