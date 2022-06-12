package kafkautils

import (
	"github.com/segmentio/kafka-go"
)

var Writer kafka.Writer

func InitWriter() {
	Writer = kafka.Writer{
		Addr:     kafka.TCP("localhost:9095"),
		Topic:    "testing-01",
		Balancer: &kafka.Hash{},
	}
}

func CloseWriter() {
	Writer.Close()
}
