package adapters

import (
	"wallet/source/application"

	"wallet/source/infrastructure/kafka"

	"github.com/gin-gonic/gin"
)

type CreateUserSchema struct {
	Email    string
	Password string
}

func CreateUserController(c *gin.Context) {
	var payload CreateUserSchema
	json_error := c.BindJSON(&payload)
	if json_error != nil {
		c.JSON(400, gin.H{
			"details": json_error.Error(),
		})
		return
	}

	message_repo := &KafkaMessageRepository{kafka.Writer}

	service := application.CreateUserService{MessageRepo: message_repo}

	service_error := service.Execute(payload.Email, payload.Password)
	if service_error != nil {
		c.JSON(500, gin.H{
			"details": "something went wrong, sorry",
		})
		return
	}
}
