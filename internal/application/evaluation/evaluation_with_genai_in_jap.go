package evaluation

import (
	"context"

	"github.com/kakky/refacgo/internal/application"
)

type EvaluationWithGenAiInJap struct {
	geminiClient application.GenAI
}

func NewEvaluationWithGenAI(geminiClient application.GenAI) *EvaluationWithGenAiInJap {
	return &EvaluationWithGenAiInJap{
		geminiClient: geminiClient,
	}
}

// インターフェースを満たすメソッドを定義
func (ev *EvaluationWithGenAiInJap) Evaluate(ctx context.Context, src []byte, filename string) (string, error) {
	return "", nil
}
