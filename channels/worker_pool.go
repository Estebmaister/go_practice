package channels

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func WorkerCreator() {
	// Create a pipeline
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 worker goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(3)
	for w := 1; w <= 3; w++ {
		go func(id int) {
			defer wg.Done()
			processJobs(ctx, id, jobs, results)
		}(w)
	}

	// Feed jobs and close when done
	go func() {
		for j := 1; j <= 9; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Start a collector
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	totalResult := 0
	for r := range results {
		totalResult += r
		fmt.Printf("Got result: %d\n", r)
	}
	fmt.Printf("Final result: %d\n", totalResult)
}

func processJobs(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return // Channel closed
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(100 * time.Millisecond) // Simulate work
			select {
			case results <- job * 2:
				// Result sent
			case <-ctx.Done():
				return
			}
		}
	}
}
