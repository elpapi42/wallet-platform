package adapters

import (
	"context"
	"encoding/json"
	"wallet/source/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageRepository struct {
	Writer kafka.Writer
}

func (r *KafkaMessageRepository) Add(message domain.Message) error {
	value, error := json.Marshal(message)
	if error != nil {
		return error
	}

	r.Writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(message.GetKey()),
		Value: value,
		Headers: []kafka.Header{
			{Key: "name", Value: []byte(message.GetName())},
		},
	})

	return nil
}
