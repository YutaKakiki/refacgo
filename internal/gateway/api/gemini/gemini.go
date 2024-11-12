package gemini

import (
	"context"
	"errors"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/kakky/refacgo/internal/config"
	"google.golang.org/api/option"
)

const (
	geminiModel = "gemini-1.5-flash"
)

type Gemini struct {
	geminiConfig config.GeminiConfig
	client       *genai.Client
	model        *genai.GenerativeModel
}

func NewGemini(geminiConfig config.GeminiConfig, ctx context.Context) *Gemini {
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiConfig.API_KEY))
	model := client.GenerativeModel(geminiModel)
	if err != nil {
		log.Fatal(err)
	}
	return &Gemini{
		geminiConfig: geminiConfig,
		client:       client,
		model:        model,
	}
}

func (gc *Gemini) Query(ctx context.Context, src []byte, prompt string) (string, error) {
	// client&modelが何らかの理由でnilの場合は早期リターン
	if gc.client == nil || gc.model == nil {
		return "", errors.New("connection to gemini failed")
	}
	// 実行が終わったらクライアントをクローズしておく
	defer gc.client.Close()
	// 受け取ったバイト配列を文字列にしたものをラップ
	code := genai.Text(string(src))
	// プロンプトをラップ
	promptText := genai.Text(prompt)
	resp, err := gc.model.GenerateContent(ctx, code, promptText)
	if err != nil {
		return "", err
	}
	var respString string
	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, part := range cand.Content.Parts {
			// Text型の場合のみレスポンス文字列に格納する
			switch p := part.(type) {
			case genai.Text:
				respString = string(p)
			}
		}
	}
	return respString, nil
}
