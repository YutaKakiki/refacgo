package evaluation

import "context"

type EvaluateByGemini struct {
	// geminiクライアント
}

func NewEvaluateByGemini() *EvaluateByGemini {
	return &EvaluateByGemini{
		// gemini: gemini,
	}
}

// インターフェースを満たすメソッドを定義
func (ev *EvaluateByGemini) Evaluate(ctx context.Context, src []byte) ([]byte, error) {
	return nil, nil
}
