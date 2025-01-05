package tests

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}
	for name, tc := range tests {
		tc := tc // rebind tc into this lexical scope
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Split(tc.input, tc.sep)

			// Option using reflect.DeepEqual with %#v
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
			}

			// Option using cmp.Diff
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("difference: %v", diff)
			}
		})
	}
}
