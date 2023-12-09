package main

import (
	"sync"
	"testing"
)

func TestSumOfSquares(t *testing.T) {
	type args struct {
		result   int
		arr      []int
		number   chan int
		response chan int
		wg       *sync.WaitGroup
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Square 1",
			args: args{1,
				[]int{1},
				make(chan int), make(chan int), &sync.WaitGroup{}},
		},
		{
			name: "Square of 5 int elements",
			args: args{55,
				[]int{1, 2, 3, 4, 5},
				make(chan int), make(chan int), &sync.WaitGroup{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer close(tt.args.number)
			defer close(tt.args.response)
			var total int
			for _, n := range tt.args.arr {
				// Launching the event receiving func and adding it to the wg
				tt.args.wg.Add(1)
				go sumOfSquares(tt.args.number, tt.args.response, tt.args.wg)
				// Sending the values to the channel funnel
				tt.args.number <- n
				// Collecting the responses
				total += <-tt.args.response
			}

			tt.args.wg.Wait()
			t.Logf("Sum of squares: %#+v\n", total)
		})
	}
}
