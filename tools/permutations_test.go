package tools

import (
	"testing"
)

func TestPerm(t *testing.T) {
	type args struct {
		a []rune
		f func([]rune)
	}
	tests := []struct {
		name string
		args args
	}{
		{"Simple slice", args{
			[]rune("abc"), func(a []rune) { println(string(a)) },
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Perm(tt.args.a, tt.args.f)
		})
	}
}
