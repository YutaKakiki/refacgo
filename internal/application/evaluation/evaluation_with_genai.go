package evaluation

import (
	"context"

	"github.com/kakky/refacgo/internal/application"
)

type EvaluationWithGenAI struct {
	geminiClient application.GenAI
}

func NewEvaluationWithGenAI(geminiClient application.GenAI) *EvaluationWithGenAI {
	return &EvaluationWithGenAI{
		geminiClient: geminiClient,
	}
}

// インターフェースを満たすメソッドを定義
func (ev *EvaluationWithGenAI) Evaluate(ctx context.Context, src []byte, filename string) (string, error) {
	return "", nil
}
