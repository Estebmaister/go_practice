// Stack Challenge
// Step #1 Complete the implementation of a Stack in Go.
// Step #2 Spawn 20 goroutines calling Push() method, this adds an item to the stack.
// Step #3 Spawn 10 goroutines calling Pop() method, this method removes the top item from the stack and send those values into a channel.
// Step #4 Read the channel from the main thread and print the values as they arrive
// Step #5 Optional - Add the method IsEmpty() that returns true if the stack is empty.
// Step #6 Optional - Add the method Peek() that returns the top item without removing it.

package channels

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Stack struct {
	mu    sync.Mutex
	items []int
}

// Push adds an element to the stack.
func (s *Stack) Push(element int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, element)
}

// Pop removes and sends the top item to the provided channel.
func (s *Stack) Pop(popResults chan int) {
	s.mu.Lock()
	val, ok := s.peekUnsafe()
	if !ok {
		s.mu.Unlock()
		return
	}
	s.items = s.items[:len(s.items)-1]
	s.mu.Unlock()
	popResults <- val
}

// RunConcurrentStack creates a Stack and 20 goroutines to stack data and
// 10 goroutines to extract data,
// managing the concurrency with channels, mutexes and waitgroups
func RunConcurrentStack() {
	stack := Stack{}
	popResults := make(chan int, 10)

	// Graceful shutdown handling
	// Closing mechanisms, chan timeout, context timeout, context cancel, and OS signals
	shutdownSignal := make(chan struct{})
	go func() {
		time.Sleep(time.Millisecond * 800)
		close(shutdownSignal) // Signal shutdown
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Spawn 20 goroutines to send data to the Stack
	for i := 0; i < 20; i++ {
		e := i
		go func(elem int) {
			stack.Push(elem)
		}(e)
	}
	// Spawn 10 goroutines to extract data to receiving channel
	for i := 0; i < 10; i++ {
		go func() {
			stack.Pop(popResults)
		}()
	}

	// Managing received data and closing mechanisms
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// Receiving and printing chan data
			case v := <-popResults:
				fmt.Printf("%v\n", v)
			case <-ctx.Done():
				fmt.Println("Context timeout, shutting down...")
				return
			case <-osSignals:
				fmt.Println("Received OS signal, shutting down...")
				return
			case <-time.After(time.Second * 2):
				fmt.Println("Timeout reached, shutting down...")
				return
			case <-shutdownSignal:
				fmt.Println("Draining remaining messages before shutdown...")
				for len(popResults) > 0 {
					fmt.Println(<-popResults)
				}
				fmt.Println("Shutdown signal received, exiting...")
				return
			}
		}
	}()

	// Wait for all goroutines to finish before exiting
	wg.Wait()
	fmt.Println("All goroutines stopped")
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items) == 0
}

// Peek returns the top item without removing it.
func (s *Stack) Peek() (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.peekUnsafe()
}

// peekUnsafe is an internal method that assumes the lock is already held.
func (s *Stack) peekUnsafe() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}
