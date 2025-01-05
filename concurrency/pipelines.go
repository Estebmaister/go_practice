package concurrency

import (
	// "math/big"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Creates a pipeline with a generator of random integers,
// filters the integers on several workers and display
// the first n primes generated.
// Using the fan in, fan out pattern
func PrimesPipeline(numberOfPrimes int) {
	CPUCount := runtime.NumCPU()
	start := time.Now() // timer

	println("\nPrimes pipeline for", numberOfPrimes, "primes with CPUs:", CPUCount)

	done := make(chan struct{})
	defer close(done) // closing all when finishing

	rand := rand.New(rand.NewSource(7))
	randNumFetcher := func() int { return rand.Intn(500_000_000) }
	randIntStream := repeatFunc(done, randNumFetcher)

	// fan out
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}
	// fan in
	fannedInStream := fanIn(done, primeFinderChannels...)

	for filteredPrime := range take(done, fannedInStream, numberOfPrimes) {
		println(filteredPrime)
	}

	println(time.Since(start))
}

// stream generator, loop call the func and send the result to
// the returned channel
func repeatFunc[T any](done <-chan struct{}, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

// controls the data taken from the stream channel
func take[T any](done <-chan struct{}, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

// couples the spanned channels per CPU to a single channel.
func fanIn[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)

	}()
	return fannedInStream
}

// Deploy a go routine to extract integers from randIntStream.
// this routine will stop when it receives the done signal or
// the receiving channel is blocked.
// It returns a channel that will receive filtered prime ints.
func primeFinder(done <-chan struct{}, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		if randomInt != 2 && randomInt%2 == 0 {
			return false
		}
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
		// efficient way avoided to test delay on execution
		// return big.NewInt(int64(randomInt)).ProbablyPrime(0)
	}

	primes := make(chan int)
	go func() {
		// when receiving the done signal close primes chan
		defer close(primes)
		// for select pattern to listen and filter primes
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}
