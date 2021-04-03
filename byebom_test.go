package byebom

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

const bomUTF8 = "\xEF\xBB\xBF"

func TestNormalize(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantStr string
	}{
		{
			name: "handle nil",
			args: args{
				r: nil,
			},
			wantNil: true,
		},
		{
			name: "trim bom when contained",
			args: args{
				r: strings.NewReader(bomUTF8 + "a"),
			},
			wantStr: "a",
		},
		{
			name: "do nothing when not contained",
			args: args{
				r: strings.NewReader("b"),
			},
			wantStr: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalize(tt.args.r)
			if tt.wantNil {
				assert.Nil(t, got)
			} else {
				gotB, err := io.ReadAll(got)
				if assert.NoError(t, err) {
					assert.Equal(t, tt.wantStr, string(gotB))
				}
			}
		})
	}
}
