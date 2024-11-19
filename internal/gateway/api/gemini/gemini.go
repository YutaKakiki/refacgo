package gemini

import (
	"context"
	"errors"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/kakkky/refacgo/internal/config"
	"google.golang.org/api/iterator"
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
func (gc *Gemini) StreamQueryResults(ctx context.Context, src []byte, prompt string, ch chan<- string) error {
	// client & modelが何らかの理由でnilの場合は早期リターン
	if gc.client == nil || gc.model == nil {
		return errors.New("connection to Gemini failed")
	}
	// 実行が終わったらクライアントをクローズしておく
	defer func() error {
		if err := gc.client.Close(); err != nil {
			return err
		}
		return nil
	}()

	// 受け取ったバイト配列を文字列にしたものをラップ
	code := genai.Text(string(src))
	// プロンプトをラップ
	promptText := genai.Text(prompt)
	// ストリーミングで逐次的に文字列を受け取れるようにする
	iter := gc.model.GenerateContentStream(ctx, promptText, code)
	ch <- "notify" //レスポンスを開始することを通知
	// 送信チャネルをクローズ
	defer close(ch)
	for {
		// ストリーミング
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		// レスポンスを文字列にキャストしてチャネルに送信
		for _, cand := range resp.Candidates {
			if cand.Content == nil {
				continue
			}
			for _, part := range cand.Content.Parts {
				// Text型の場合のみレスポンス文字列に格納する
				switch p := part.(type) {
				case genai.Text:
					ch <- string(p)
				}
			}
		}
	}

	return nil
}

func (gc *Gemini) QueryResuluts(ctx context.Context, src []byte, prompt string, ch chan<- string) error {
	// client & modelが何らかの理由でnilの場合は早期リターン
	if gc.client == nil || gc.model == nil {
		return errors.New("connection to Gemini failed")
	}
	// 実行が終わったらクライアントをクローズしておく
	defer func() error {
		if err := gc.client.Close(); err != nil {
			return err
		}
		return nil
	}()

	// 受け取ったバイト配列を文字列にしたものをラップ
	code := genai.Text(string(src))
	// プロンプトをラップ
	promptText := genai.Text(prompt)
	resp, err := gc.model.GenerateContent(ctx, promptText, code)
	if err != nil {
		return err
	}
	// 送信チャネルをクローズ
	defer close(ch)
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				// Text型の場合のみレスポンス文字列に格納する
				switch p := part.(type) {
				case genai.Text:
					ch <- string(p)
				}
			}
		}
	}
	return nil
}
