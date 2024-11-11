package evaluation

import "context"

type Evaluate interface {
	Evaluate(ctx context.Context, src []byte) ([]byte, error)
}
