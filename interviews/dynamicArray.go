package interviews

import (
	"fmt"
)

func dynamicArray(n int32, queries [][]int32) []int32 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}() // Panic catcher

	var results []int32
	var lastAnswer int32 = 0
	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arr = append(arr, []int32{})
	}

	for _, query := range queries {
		x, y := query[1], query[2]
		arrN := (x ^ lastAnswer) % n
		if query[0] == 1 {
			arr[arrN] = append(arr[arrN], y)
		} else {
			lastAnswer = arr[arrN][int(y)%len(arr[arrN])]
			results = append(results, lastAnswer)
			fmt.Println(lastAnswer)
		}
	}

	return results
}
