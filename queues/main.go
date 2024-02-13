package main

import (
	"os"
	"time"
)

func main() {
	// Get topic names from environment variables
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		panic("KAFKA_TOPIC environment variable not set")
	}

	addr := "localhost:9092"
	// Launch producer
	go producer(addr, topic, []string{"second", "attempt"})

	// Launch consumer and close it after timeout
	done := make(chan struct{})
	go consumer(addr, topic, done)
	time.Sleep(20 * time.Second)
	close(done)
}
