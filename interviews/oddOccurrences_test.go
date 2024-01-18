package interviews

import "testing"

func TestOddOccurrences(t *testing.T) {
	type args struct {
		oddArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"7 elements with a 7 odd",
			args{
				[]int{7, 1, 2, 3, 1, 2, 3},
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddOccurrences(tt.args.oddArray); got != tt.want {
				t.Errorf("oddOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
