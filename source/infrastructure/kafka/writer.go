package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func InitWriter() {
	Writer = &kafka.Writer{
		Addr:      kafka.TCP("localhost:9095"),
		Topic:     "testing-01",
		Balancer:  &kafka.Hash{},
		BatchSize: 1,
	}
}

func CloseWriter() {
	log.Println("closing kafka-go writer")
	Writer.Close()
	log.Println("kafka-go writer closed")
}
