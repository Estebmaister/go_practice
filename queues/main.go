package main

import (
	"os"
)

func main() {
	// Get topic names from environment variables
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		panic("KAFKA_TOPIC environment variable not set")
	}

	addr := "localhost:9092"

	go producer(addr, topic, []string{"second", "attempt"})

	consumer(addr, topic)
}
