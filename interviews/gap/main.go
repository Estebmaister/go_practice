package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	numbers := []int{3, 5, -9, 1, -3, 6}
	fmt.Println(calculateSums(numbers))
}

// Problem Statement: Concurrent Sum Calculation
//
// Given an input slice of integers, divide it into two halves. Then, concurrently calculate the sum of the elements of each half. Finally, display the individual sums along with the sum of all elements.
//
// Input => [3, 5, -9, 1, -3, 6]
// Output => -1, 4, 3
func calculateSums(numbers []int) []int {
	var leftSum, rightSum atomic.Int64
	mid := len(numbers) / 2
	leftPart, rightPart := numbers[:mid], numbers[mid:]

	var wg sync.WaitGroup
	wg.Add(2)

	go sumSlice(&wg, leftPart, &leftSum)
	go sumSlice(&wg, rightPart, &rightSum)

	wg.Wait()

	totalSum := leftSum.Load() + rightSum.Load()
	return []int{int(leftSum.Load()), int(rightSum.Load()), int(totalSum)}
}

func sumSlice(wg *sync.WaitGroup, slice []int, sum *atomic.Int64) {
	defer wg.Done()
	var localSum int64

	for _, num := range slice {
		localSum += int64(num) // Accumulate sum locally
	}

	sum.Add(localSum) // Atomic operation to avoid race conditions
}
