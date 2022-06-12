package main

import (
	"fmt"

	"wallet/source/adapters"
	"wallet/source/application"
	"wallet/source/infrastructure/kafkautils"
)

func main() {
	kafkautils.InitWriter()

	for i := 0; i <= 5; i = i + 1 {
		repo := &adapters.KafkaMessageRepository{
			Writer: kafkautils.Writer,
		}

		service := application.CreateUserService{
			MessageRepo: repo,
		}

		error := service.Execute("hola@email.com", "123456")
		if error != nil {
			fmt.Println(error)
		}

		fmt.Println("Message sent")
	}

	kafkautils.CloseWriter()
}
