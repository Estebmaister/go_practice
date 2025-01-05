package concurrency

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"testing"
)

func TestNewSema(t *testing.T) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	s := NewIO(5, l)
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			s.LimitedProcess(fmt.Sprintf("file-%d", x))
		}(i)
	}
	wg.Wait()
}
