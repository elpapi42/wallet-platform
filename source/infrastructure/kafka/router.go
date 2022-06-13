package kafka

import (
	"wallet/source/adapters/handlers"

	"github.com/segmentio/kafka-go"
)

func getHandlersForMessage(message kafka.Message) []func(kafka.Message) error {
	handlerMap := map[string][]func(kafka.Message) error{
		"CreateUser": {handlers.SaveUserHandler},
	}

	var messageName string
	for _, element := range message.Headers {
		if element.Key == "name" {
			messageName = string(element.Value)
			break
		}
	}

	return handlerMap[messageName]
}
