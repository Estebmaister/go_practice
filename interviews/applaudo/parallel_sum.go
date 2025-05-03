package main

import (
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	workers := 3
	println(ParallelSum(numbers, workers)) // Output esperado: 55
}

func ParallelSum(numbers []int, workers int) int {
	elementsPerWorker := len(numbers) / workers
	wg := &sync.WaitGroup{}
	resultChan := make(chan int, workers)

	for i := range workers {
		initPoint := i * elementsPerWorker
		endPoint := initPoint + elementsPerWorker
		if i == workers-1 {
			endPoint += len(numbers) % workers
		}
		wg.Add(1)
		go sumWorker(numbers[initPoint:endPoint], resultChan, wg)
	}

	wg.Wait()
	var totalSum int
	for range workers {
		totalSum += <-resultChan
	}
	close(resultChan)

	return totalSum
}

func sumWorker(numbers []int, returnChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	returnChan <- sum
}
