package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func producer(addr, topic string, msgs []string) {
	partition := 0

	// Create a producer
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// Set the deadline
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	// Transform the messages
	messages := make([]kafka.Message, len(msgs))
	for idx, v := range msgs {
		messages = append(messages, kafka.Message{
			Key: []byte(strconv.Itoa(idx)), Value: []byte(v)})
	}

	// Send messages
	_, err = conn.WriteMessages(messages...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	} else {
		log.Println("Messages sent to", topic)

	}

	// Close producer
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
