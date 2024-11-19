package eval

import (
	"context"
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakky/refacgo/internal/domain"
	"github.com/kakky/refacgo/internal/presenter"
	"github.com/kakky/refacgo/internal/presenter/indicater"
	"github.com/urfave/cli/v2"
	"go.uber.org/mock/gomock"
)

//go:embed testdata/prompt/with_genai_prompt.txt
var expectedPrompt []byte

//go:embed testdata/prompt/with_genai_in_jap_prompt.txt
var expectedPromptInJap []byte

func TestEvalCmd(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockGenAI := domain.NewMockGenAI(ctrl)
	mockEvalPrinter := presenter.NewMockEvalPrinter(ctrl)
	mockIndicater := indicater.NewMockIndicater(ctrl)
	srcArg := []byte("This is sample code.\n")
	respString := []string{"This is comments of evalutated code!!!", "This is response from Mock!!!"}
	respStringInJap := []string{"とてもいいコードです！！", "とてもいいテストコードです！！"}
	srcArgWithDesc := []byte("これはテストで用いるためのものです。 :\n\n\nThis is sample code.\n")
	// チャネルから受信した文字列を格納する配列
	var got string
	tests := []struct {
		name     string
		args     []string
		mockFunc func()
		want     string
		wantErr  bool
	}{
		{
			name: "フラグなしでコマンドを叩くと評価コメントが返る",
			args: []string{"refacgo", "eval", "./testdata/src/sample.txt"},
			mockFunc: func() {

				mockIndicater.EXPECT().Start()
				mockGenAI.EXPECT().StreamQueryResults(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, src []byte, prompt string, ch chan<- string) {
						// 正確にプロンプト・ソースコードをQueryに渡しているか
						if diff := cmp.Diff(string(expectedPrompt), prompt); diff != "" {
							t.Errorf("prompt received from EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						if diff := cmp.Diff(srcArg, src); diff != "" {
							t.Errorf("src received from  mismatch EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						// 文字列をチャネルに送信
						ch <- "notify"
						for _, rs := range respString {
							ch <- rs
						}
						close(ch)
					},
				)
				mockIndicater.EXPECT().Stop()
				mockEvalPrinter.EXPECT().Print(gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, ch <-chan string) {
						for text := range ch {
							got += text
						}
					},
				)

			},
			want:    respString[0] + respString[1],
			wantErr: false,
		},
		{
			name: "-jフラグをつけてコマンドを叩くと日本語による評価コメントが返る",
			args: []string{"refacgo", "eval", "-j", "./testdata/src/sample.txt"},
			mockFunc: func() {
				mockIndicater.EXPECT().Start()
				mockGenAI.EXPECT().StreamQueryResults(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, src []byte, prompt string, ch chan<- string) {
						// 正確にプロンプト・ソースコードをQueryに渡しているか
						if diff := cmp.Diff(string(expectedPromptInJap), prompt); diff != "" {
							t.Errorf("prompt received from EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						if diff := cmp.Diff(src, src); diff != "" {
							t.Errorf("src received from  mismatch EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						defer close(ch)
						// 文字列をチャネルに送信
						ch <- "notify"
						for _, rs := range respStringInJap {
							ch <- rs
						}
					},
				)
				mockIndicater.EXPECT().Stop()
				mockEvalPrinter.EXPECT().Print(gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, ch <-chan string) {
						for text := range ch {
							got += text
						}
					},
				)
			},
			want:    respStringInJap[0] + respStringInJap[1],
			wantErr: false,
		},
		{
			name: "-jフラグをつけ、-descフラグをつけてコマンドを叩くと日本語による評価コメントが返り、ソースコードに説明が追加される",
			args: []string{"refacgo", "eval", "-j", "-desc", "これはテストで用いるためのものです。", "./testdata/src/sample.txt"},
			mockFunc: func() {
				mockIndicater.EXPECT().Start()
				mockGenAI.EXPECT().StreamQueryResults(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, src []byte, prompt string, ch chan<- string) {
						// 正確にプロンプト・ソースコードをQueryに渡しているか
						if diff := cmp.Diff(string(expectedPromptInJap), prompt); diff != "" {
							t.Errorf("prompt received from EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						if diff := cmp.Diff(srcArgWithDesc, src); diff != "" {
							t.Errorf("src received from  mismatch EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						defer close(ch)
						// 文字列をチャネルに送信
						ch <- "notify"
						for _, rs := range respStringInJap {
							ch <- rs
						}
					},
				)
				mockEvalPrinter.EXPECT().Print(gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, ch <-chan string) {
						for text := range ch {
							got += text
						}
					},
				)
				mockIndicater.EXPECT().Stop()
			},
			want:    respStringInJap[0] + respStringInJap[1],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Cleanup(func() {
				got = ""
			})
			tt.mockFunc()
			ctx := context.Background()
			app := &cli.App{
				Name:        "refacgo",
				Description: "A Go-based command-line tool that evaluates the code in a specified Go file and provides refactoring suggestions powered by AI",
				Commands: []*cli.Command{
					EvalCmd(ctx, mockGenAI, mockEvalPrinter, mockIndicater),
				},
			}
			if err := app.RunContext(ctx, tt.args); err != nil {
				t.Errorf("Error in Running CLI : %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("expected output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
