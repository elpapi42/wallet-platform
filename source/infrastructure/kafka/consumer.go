package kafka

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
	lock   *sync.Mutex
	stop   bool
}

func NewKafkaConsumer(brokers []string, topic string, group string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: group,
		}),
		lock: &sync.Mutex{},
	}
}

func (k *KafkaConsumer) Start() {
	k.stop = false
	k.lock.Lock()

	go k.consume()
}

func (k *KafkaConsumer) consume() {
	defer k.lock.Unlock()

	for !k.stop {
		k.iterate()
	}
}

func (k *KafkaConsumer) iterate() {
	fetchCtx, fetchCancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer fetchCancel()
	m, fetchError := k.reader.FetchMessage(fetchCtx)
	if fetchError != nil && errors.Is(fetchError, context.DeadlineExceeded) {
		return
	} else if fetchError != nil {
		log.Println("failed to fetch message:", fetchError)
		return
	}

	log.Println("message received:", string(m.Key), string(m.Value))

	handlers := getHandlersForMessage(m)

	for _, handler := range handlers {
		handler(m)
	}

	commitCtx, commitCancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer commitCancel()
	commitError := k.reader.CommitMessages(commitCtx, m)
	if commitError != nil {
		log.Fatalln("failed to commit message:", commitError)
		return
	}
}

func (k *KafkaConsumer) Stop() {
	log.Println("stopping kafka consumer")

	k.stop = true
	defer k.lock.Unlock()
	k.lock.Lock()

	log.Println("kafka consumer stopped")
}

func (k *KafkaConsumer) Close() {
	log.Println("closing kafka consumer")

	k.Stop()
	k.reader.Close()

	log.Println("kafka consumer closed")
}
