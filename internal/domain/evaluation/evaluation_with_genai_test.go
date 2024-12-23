package evaluation

import (
	"context"
	_ "embed"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/refacgo/internal/domain"
	"go.uber.org/mock/gomock"
)

//go:embed testdata/prompt/with_genai_prompt.txt
var expectedPrompt []byte

func TestEvauationWithGenAI(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockGenAI := domain.NewMockGenAI(ctrl)
	srcArg := []byte("This is sample code.")
	respString := []string{"This is comments of evalutated code!!!", "This is response from Mock!!!"}
	type args struct {
		src      []byte
		filename string
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		want     string
	}{
		{
			name: "GenAIにソースコード・プロンプトを正常に渡し、非同期的にチャネルに文字列を送信できる",
			mockFunc: func() {
				mockGenAI.EXPECT().StreamQueryResults(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(
					func(ctx context.Context, src []byte, prompt string, ch chan<- string) {
						// 正確にプロンプト・ソースコードをQueryに渡しているか
						if diff := cmp.Diff(string(expectedPrompt), prompt); diff != "" {
							t.Errorf("prompt received from EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						if diff := cmp.Diff(srcArg, src); diff != "" {
							t.Errorf("src received from  mismatch EvaluationWithGenAI mismatch (-want +got):\n%s", diff)
						}
						defer close(ch)
						// 文字列をチャネルに送信
						for _, rs := range respString {
							ch <- rs
						}
					},
				)
			},
			args: args{
				src:      srcArg,
				filename: "test.go",
			},
			want: respString[0] + respString[1],
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			ctx := context.Background()
			evaluation := NewEvaluationWithGenAI(mockGenAI)
			ch := make(chan string)
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := evaluation.Evaluate(ctx, tt.args.src, tt.args.filename, ch); err != nil {
					t.Error(err)
				}
			}()
			// チャネルから文字列を受信
			var ss []string
			for text := range ch {
				ss = append(ss, text)
			}
			got := ss[0] + ss[1]
			if got != tt.want {
				t.Errorf("evaluated response not match,want: %q,got: %q", tt.want, got)
			}
			wg.Wait()
		})
	}
}
