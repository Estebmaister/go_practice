package interviews

import (
	"sync"
	"time"
)

func sumOfSquares(number chan int, response chan int, wg *sync.WaitGroup) {
	// Signaling end of waiting
	defer wg.Done()
	sum1 := <-number
	sum1 = sum1 * sum1

	response <- sum1
	time.Sleep(time.Second)
}
