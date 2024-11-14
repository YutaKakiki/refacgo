package application

import "context"

//go:generate mockgen -package=application -source=./genai_interface.go -destination=./genai_mock.go
type GenAI interface {
	Query(ctx context.Context, src []byte, prompt string, ch chan<- string) error
}
