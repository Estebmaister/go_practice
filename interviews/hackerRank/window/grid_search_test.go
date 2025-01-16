package window

import "testing"

func Test_gridPatternSearch(t *testing.T) {
	type args struct {
		G []string
		P []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "fail",
			args: args{G: []string{
				"7283455864",
				"6731158619",
				"8982742668",
				"3830841210",
				"1224754986"}, P: []string{"1234", "4321", "9999"}},
			want: NO,
		}, {
			name: "border case, two adjacent pattern find",
			args: args{G: []string{
				"7283455864",
				"6731158619",
				"8989505643",
				"2229505813",
				"5633845374",
				"6473530293",
				"7053106601",
				"0834282956",
				"4607924137"}, P: []string{"9505", "3845", "3530"}},
			want: YES,
		}, {
			name: "border case, grid and pattern length equal",
			args: args{G: []string{
				"11111",
				"11111",
				"11111",
				"11111"}, P: []string{"11111", "11111", "11111"}},
			want: YES,
		}, {
			name: "too small arrays",
			args: args{G: []string{
				"1"}, P: []string{""}},
			want: NO,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridPatternSearch(tt.args.G, tt.args.P); got != tt.want {
				t.Errorf("gridPatternSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
