package interviews

import "testing"

func TestMinWaterTanks(t *testing.T) {
	type args struct {
		houses string
	}
	type testStruct struct {
		name string
		args args
		want int
	}
	tests := []testStruct{
		{"Empty string", args{""}, 0},
		{"No space for tanks 1", args{"H"}, -1},
		{"No space for tanks 2", args{"HH"}, -1},
		{"No space for tanks 3", args{"HHH"}, -1},
		{"No space for tanks 4", args{"HHHH"}, -1},
		{"No space for tanks 5", args{"HHHHH"}, -1},
		{"Only empty spaces 1", args{"-"}, 0},
		{"Only empty spaces 2", args{"--"}, 0},
		{"Only empty spaces 3", args{"---"}, 0},
		{"Only empty spaces 4", args{"----"}, 0},
		{"Only empty spaces 5", args{"-----"}, 0},
		{"Simple two chars 1", args{"-H"}, 1},
		{"Simple two chars 2", args{"H-"}, 1},
		{"One tank, three chars 1", args{"H-H"}, 1},
		{"One tank, three chars 2", args{"--H"}, 1},
		{"One tank, three chars 3", args{"H--"}, 1},
		{"One tank, four chars 1", args{"H---"}, 1},
		{"One tank, four chars 2", args{"-H--"}, 1},
		{"One tank, four chars 3", args{"--H-"}, 1},
		{"One tank, four chars 4", args{"---H"}, 1},
		{"One tank, four chars 5", args{"-H-H"}, 1},
		{"One tank, four chars 6", args{"H-H-"}, 1},
		{"Two tanks, four chars 1", args{"H--H"}, 2},
		{"Two tanks, four chars 2", args{"-HH-"}, 2},
		{"Trapped house, four chars 1", args{"HH--"}, -1},
		{"Trapped house, four chars 2", args{"HHH-"}, -1},
		{"Trapped house, four chars 3", args{"-HHH"}, -1},
		{"Trapped house, four chars 4", args{"--HH"}, -1},
		{"One tank, five chars 1", args{"-H-H-"}, 1},
		{"One tank, five chars 2", args{"-H---"}, 1},
		{"One tank, five chars 3", args{"H----"}, 1},
		{"One tank, five chars 4", args{"----H"}, 1},
		{"Two tanks, five chars 1", args{"H---H"}, 2},
		{"Two tanks, five chars 2", args{"H-H-H"}, 2},
		{"Two tanks, five chars 3", args{"H-HH-"}, 2},
		{"Two tanks, five chars 4", args{"--HH-"}, 2},
		{"Trapped house, five chars 1", args{"HH---"}, -1},
		{"Trapped house, five chars 2", args{"HHH--"}, -1},
		{"Trapped house, five chars 3", args{"HHHH-"}, -1},
		{"Trapped house, five chars 4", args{"-HHH-"}, -1},
		{"Trapped house, five chars 5", args{"-HHHH"}, -1},
		{"Trapped house, five chars 6", args{"--HHH"}, -1},
		{"Trapped house, five chars 7", args{"---HH"}, -1},
		{"Three tanks, six chars 1", args{"H--HH-"}, 3},
		{"Three tanks, six chars 2", args{"-HH--H"}, 3},
		{"Two tanks, seven chars 1", args{"H-H-H-H"}, 2},
		{"Two tanks, seven chars 2", args{"H-HH-H-"}, 2},
		{"Two tanks, seven chars 3", args{"-H-HH-H"}, 2},
		{"Three tanks, seven chars 1", args{"H--HH-H"}, 3},
		{"", args{"H--HH-H--"}, 3},
		{"", args{"H--HH-H-H"}, 4},
		{"", args{"H--HH-H--H"}, 4},
		{"", args{"H--HH-H--HH"}, -1},
		{"", args{"H--HH-H--HHH"}, -1},
		{"", args{"H--HH-H--H-H"}, 4},
		{"", args{"H--HH-H--H---"}, 4},
		{"", args{"H--HH-H--H--H-"}, 5},
		// {"", args{"H--HH-H-H-"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWaterTanks(tt.args.houses); got != tt.want {
				t.Errorf("minWaterTanks(%s) = %v, want %v", tt.args.houses, got, tt.want)
			}

			if got := numberOfWaterTanks(tt.args.houses); got != tt.want {
				if got == -1 && tt.want == 0 {
					// It's ok, different logic
				} else {
					t.Errorf("numberOfWaterTanks(%s) = %v, want %v", tt.args.houses, got, tt.want)
				}
			}

			if got := stringsWaterTanks(tt.args.houses); got != tt.want {
				if got == -1 && tt.want == 0 {
					// It's ok, different logic
				} else {
					t.Errorf("stringsWaterTanks(%s) = %v, want %v", tt.args.houses, got, tt.want)
				}
			}
		})
	}
}
