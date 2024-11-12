package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddDescToSrc(t *testing.T) {
	t.Parallel()
	type args struct {
		src  []byte
		desc string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "ソースファイルの先頭に改行を伴って説明文が追加される",
			args: args{
				src:  []byte("This is source."),
				desc: "This is description for the source.",
			},
			want: []byte("This is description for the source. :\n\n\nThis is source."),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := AddDescToSrc(tt.args.src, tt.args.desc)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("AddDescToSrc() return byte mismatch(-want +got):\n%s", diff)
			}
		})
	}
}
