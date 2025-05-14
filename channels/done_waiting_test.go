package channels

import (
	"testing"
	"time"
)

// DoneWaiting_Test demonstrates how to wait for the first signal from multiple channels.
func Benchmark_DoneWaiting(b *testing.B) {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	b.Logf("Done after %v\n", time.Since(start))
}

// sig returns a channel that will send a signal after the specified duration.
// It is used to simulate a signal that can be received after a certain time.
func sig(after time.Duration) <-chan any {
	c := make(chan any)
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// or waits for the first of the given channels to send a signal and returns that channel.
// It is used to combine multiple channels into one, returning the first one that receives a signal.
// This is a simplified version of the or function that handles multiple channels.
// It uses a recursive approach to handle more than two channels.
// The function will block until one of the channels sends a signal, at which point it will return that channel.
// If no channels are provided, it returns nil.
func or(channels ...<-chan any) <-chan any {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan any)
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
