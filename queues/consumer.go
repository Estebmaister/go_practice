package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func consumer(addr, topic string) {

	// Create a consumer
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{addr},
		GroupID:  "my-group",
		Topic:    topic,
		MinBytes: 10e3, // Minimum read batch size
		MaxBytes: 10e6, // Maximum read batch size
	})

	// Read messages
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(context.TODO())
		if err == nil {
			fmt.Println("Received message from", topic, ":", string(msg.Value))
		} else {
			fmt.Println(err)
		}
	}
}
