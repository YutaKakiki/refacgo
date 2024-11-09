package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoadFile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		filepath string
		want     []byte
		wantErr  bool
	}{
		{
			name:     "ファイルを読み込み、正しいバイトスライスを返す",
			filepath: "./testdata/sample.txt",
			want:     []byte("This is Sample File.\n"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := LoadFile(tt.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("utils.LoadFile() error = %v ,wantErr = %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("loadFile return byte mismatch(-want +got):\n%s", diff)
			}
		})
	}
}
