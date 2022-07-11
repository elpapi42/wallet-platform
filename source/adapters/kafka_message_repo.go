package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"wallet/source/domain/messages"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageRepository struct {
	Writer *kafka.Writer
}

func (r *KafkaMessageRepository) Add(messages []messages.Message) error {
	kafka_messages := make([]kafka.Message, len(messages))

	// Parse the aggregate events into messages
	for index, message := range messages {
		value, error := json.Marshal(message)
		if error != nil {
			return error
		}

		kafka_message := kafka.Message{
			Key:   []byte(message.GetKey()),
			Value: value,
			Headers: []kafka.Header{
				{Key: "name", Value: []byte(message.GetName())},
			},
		}

		log.Println("message created:", string(kafka_message.Key), string(kafka_message.Value))

		kafka_messages[index] = kafka_message
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := r.Writer.WriteMessages(ctx, kafka_messages...)
	if err != nil {
		log.Println("failed to write message to kafka:", err)
		return err
	}

	log.Println("messages sent")

	return nil
}
