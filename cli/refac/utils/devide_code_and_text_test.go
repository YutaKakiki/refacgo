package utils

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed testdata/mix_of_code_and_text.txt
var mix string

//go:embed testdata/only_text.txt
var onlyText string

//go:embed testdata/devided_code.txt
var devidedCode string

//go:embed testdata/devided_text.txt
var devidedText string

func TestDecideCodeAndText(t *testing.T) {
	t.Parallel()
	type want struct {
		code string
		text string
	}
	tests := []struct {
		name    string
		arg     string
		want    want
		wantErr bool
	}{
		{
			name: "引数にとった文字列をコードとテキストに分離して返す",
			arg:  mix,
			want: want{
				code: devidedCode,
				text: devidedText,
			},
			wantErr: false,
		},
		{
			name:    "引数にとった文字列にコードブロックがない場合、エラーをかえす",
			arg:     onlyText,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			code, text, err := DevideCodeAndText(tt.arg)
			fmt.Println(code)
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(tt.want.code, code); diff != "" {
				t.Errorf("DevideCodeAndText() return code mismatch (-want +got): %s", diff)
			}
			if diff := cmp.Diff(tt.want.text, text); diff != "" {
				t.Errorf("DevideCodeAndText() return text mismatch (-want +got): %s", diff)
			}

		})
	}
}
