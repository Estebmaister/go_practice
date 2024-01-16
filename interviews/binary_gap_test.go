package main

import "testing"

func Test_BinaryGap(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"33 binary 100001",
			args{33},
			4,
		},
		{
			"92 binary 1011100",
			args{92},
			1,
		},
		{
			"0 binary 0",
			args{0},
			0,
		},
		{
			"1 binary 1",
			args{1},
			0,
		},
		{
			"100000 binary 11000011010100000",
			args{100000},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.n); got != tt.want {
				t.Errorf("binary_gap() = %v, want %v", got, tt.want)
			}
		})
	}
}
