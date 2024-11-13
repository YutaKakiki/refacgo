package evaluation

import "context"

type Evaluation interface {
	Evaluate(ctx context.Context, src []byte, filename string, ch chan<- string) error
}
