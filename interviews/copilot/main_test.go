package main

import "testing"

func TestRPNcalculator(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			"Success",
			args{"3 4 +"},
			7,
			false,
		}, {
			"Success",
			args{"10 4 -"},
			6,
			false,
		}, {
			"Success",
			args{"6 7 *"},
			42,
			false,
		}, {
			"Success",
			args{"8 2 /"},
			4,
			false,
		}, {
			"Success",
			args{"5 1 2 + 4 * + 3 -"},
			14,
			false,
		}, {
			"Success",
			args{"9 sqrt 3 /"},
			1,
			false,
		}, {
			"Success",
			args{"9 sqrt"},
			3,
			false,
		}, {
			"Failure",
			args{"+"},
			0,
			true,
		}, {
			"Failure",
			args{"3 -"},
			0,
			true,
		}, {
			"Failure",
			args{"10 0 /"},
			0,
			true,
		}, {
			"Failure",
			args{"3 3"},
			0,
			true,
		}, {
			"Failure",
			args{"-4 sqrt"},
			0,
			true,
		}, {
			"Failure",
			args{"sqrt"},
			0,
			true,
		}, {
			"Failure",
			args{"abc"},
			0,
			true,
		}, {
			"Failure",
			args{"4 5 &"},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RPNcalculator(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPNcalculator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RPNcalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
