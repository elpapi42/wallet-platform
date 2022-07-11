package kafka

import (
	"github.com/segmentio/kafka-go"
)

func getHandlersForMessage(message kafka.Message) []func(kafka.Message) error {
	handlerMap := map[string][]func(kafka.Message) error{
		//"DepositReceived": {handlers.UpdateWalletBalanceByDepositHandler},
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
