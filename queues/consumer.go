package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func consumer(addr, topic string, doneChan <-chan struct{}) {

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
		select {
		case <-doneChan:
			log.Println("Consumer closed")
			return
		default:
			msg, err := reader.ReadMessage(context.TODO())
			if err == nil {
				log.Println("Received message from", topic, ":", string(msg.Value))
			} else {
				log.Println(err)
			}
		}
	}
}
