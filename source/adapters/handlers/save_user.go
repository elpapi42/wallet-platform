package handlers

import (
	"encoding/json"
	"log"

	"wallet/source/application"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type SaveUserSchema struct {
	Id       uuid.UUID
	Email    string
	Password string
}

func SaveUserHandler(message kafka.Message) error {
	log.Println("processing message:", string(message.Key), string(message.Value))

	var schema SaveUserSchema

	jsonError := json.Unmarshal(message.Value, &schema)
	if jsonError != nil {
		return jsonError
	}

	service := application.SaveUserService{}

	serviceError := service.Execute(schema.Id, schema.Email, schema.Password)
	if serviceError != nil {
		return serviceError
	}

	log.Println("message processed:", string(message.Key), string(message.Value))

	return nil
}
