package byebom

import (
	"io"
	"strings"
	"testing"

	"golang.org/x/text/encoding"

	"github.com/stretchr/testify/assert"
)

const utf8BOM = "\ufeff"

func TestNewUTF8Reader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name               string
		args               args
		want               string
		wantInvalidUTF8Err bool
	}{
		{
			name: "with BOM",
			args: args{
				r: strings.NewReader(utf8BOM + "a"),
			},
			want: "a",
		},
		{
			name: "without BOM",
			args: args{
				r: strings.NewReader("a"),
			},
			want: "a",
		},
		{
			name: "not UTF-8(euc-jp)",
			args: args{
				r: strings.NewReader("\xa5\xa8\xa5\xe9\xa1\xbc"), // means "エラー"
			},
			wantInvalidUTF8Err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUTF8Reader(tt.args.r)
			gotB, err := io.ReadAll(got)
			if tt.wantInvalidUTF8Err {
				assert.Equal(t, encoding.ErrInvalidUTF8, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, string(gotB))
			}
		})
	}
}
