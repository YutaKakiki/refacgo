package application

import "context"

type GenAI interface {
	Query(ctx context.Context, src []byte, prompt string) (string, error)
}
