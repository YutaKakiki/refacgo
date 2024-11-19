package refactoring

import (
	"context"
	_ "embed"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakky/refacgo/internal/domain"
	"go.uber.org/mock/gomock"
)

//go:embed testdata/prompt/with_genai_prompt_in_jap.txt
var expectedPromptInJap string

func TestRefactoringWithGenAiInJap(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockGenAI := domain.NewMockGenAI(ctrl)
	srcArg := []byte("This is sample code.")
	respString := []string{"これはリファクタリングされたコードです！", "これはモックからのレスポンスです！"}
	type arg struct {
		src      []byte
		filename string
	}
	tests := []struct {
		name     string
		arg      arg
		mockFunc func()
		want     string
		wantErr  bool
	}{
		{
			name: "GenAIにプロンプト・ソースコードを正常渡し、非同期にリファクタリング結果を得る",
			arg: arg{
				src:      srcArg,
				filename: "test.go",
			},
			mockFunc: func() {
				mockGenAI.EXPECT().QueryResuluts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, src []byte, prompt string, ch chan<- string) error {
						if diff := cmp.Diff(expectedPromptInJap, prompt); diff != "" {
							t.Errorf("prompt received from RefactoringWithGenAiInJap mismatch (-want +got):\n%s", diff)
						}
						if diff := cmp.Diff(srcArg, src); diff != "" {
							t.Errorf("src received from  mismatch RefactoringWithGenAiInJap mismatch (-want +got):\n%s", diff)
						}
						defer close(ch)
						for _, resp := range respString {
							ch <- resp
						}
						return nil
					})
			},
			want:    respString[0] + respString[1],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.mockFunc()
			refactoring := NewRefactoringWithGenAiInJap(mockGenAI)
			ctx := context.Background()
			ch := make(chan string)
			var wg sync.WaitGroup
			wg.Add(1)
			go func() error {
				defer wg.Done()
				if err := refactoring.Refactor(ctx, tt.arg.src, tt.arg.filename, ch); err != nil {
					return err
				}
				return nil
			}()
			var got string
			for s := range ch {
				got += s
			}
			if got != tt.want {
				t.Errorf("evaluated response not match,want: %q,got: %q", tt.want, got)
			}
			wg.Wait()
		})
	}

}
