package strings

import "testing"

func TestMergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Same length words",
			args: args{word1: "abc", word2: "pqr"},
			want: "apbqcr",
		},
		{
			name: "Different length words",
			args: args{word1: "abcd", word2: "pq"},
			want: "apbqcd",
		},
		{
			name: "Empty word",
			args: args{word1: "", word2: "pq"},
			want: "pq",
		},
		{
			name: "Empty word",
			args: args{word1: "abc", word2: ""},
			want: "abc",
		},
		{
			name: "Empty strings",
			args: args{word1: "", word2: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("MergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
