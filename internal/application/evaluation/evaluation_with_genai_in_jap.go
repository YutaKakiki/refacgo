package evaluation

import (
	"context"
	"fmt"
	"path/filepath"

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

// インターフェースを満たすメソッドを定義
func (ev *EvaluationWithGenAiInJap) Evaluate(ctx context.Context, src []byte, filename string) (string, error) {
	path := filepath.Join("internal", "application", "evaluation", "genai_instruction_in_jap.txt")
	instruction, err := loadfile.LoadFile(path)
	if err != nil {
		panic(err)
	}
	prompt := fmt.Sprintf("このファイルの名前は%qです。\n\n%v\n\n", filename, string(instruction))
	resp, err := ev.genAI.Query(ctx, src, prompt)
	if err != nil {
		return "", err
	}
	return resp, nil
}
