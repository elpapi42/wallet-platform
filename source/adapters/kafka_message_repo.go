package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"wallet/source/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageRepository struct {
	Writer *kafka.Writer
}

func (r *KafkaMessageRepository) Add(message domain.Message) error {
	value, error := json.Marshal(message)
	if error != nil {
		return error
	}

	m := kafka.Message{
		Key:   []byte(message.GetKey()),
		Value: value,
		Headers: []kafka.Header{
			{Key: "name", Value: []byte(message.GetName())},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := r.Writer.WriteMessages(ctx, m)
	if err != nil {
		log.Println("failed to write message to kafka: ", err)
		return err
	}

	log.Println("message sent:", string(m.Key), string(m.Value))

	return nil
}
