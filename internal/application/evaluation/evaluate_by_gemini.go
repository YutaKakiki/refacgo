package evaluation

import "github.com/kakky/refacgo/internal/application"

type EvaluateByGemini struct {
	// gemini APIインターフェース
	gemini application.Gemini
}

func NewEvaluation(gemini application.Gemini) *EvaluateByGemini {
	return &EvaluateByGemini{
		// gemini: gemini,
	}
}

func (ev *EvaluateByGemini) Evaluate(filepath string) ([]byte, error) {
	return nil, nil
}

func (ev *EvaluateByGemini) EvaluateInJapanese(filepath string) {

}
