package controllers

import (
	"wallet/source/adapters"
	"wallet/source/application"
	"wallet/source/infrastructure/eventstore"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateWalletSchema struct {
	User     uuid.UUID
	Currency string
}

func CreateWalletController(c *gin.Context) {
	var payload CreateWalletSchema
	json_error := c.BindJSON(&payload)
	if json_error != nil {
		c.JSON(400, gin.H{
			"details": json_error.Error(),
		})
		return
	}

	wallet_repo := &adapters.EventstoreWalletRepository{EventstoreClient: eventstore.GetClient()}

	service := application.CreateWalletService{WalletRepository: wallet_repo}

	wallet, service_error := service.Execute(payload.User, payload.Currency)
	if service_error != nil {
		c.JSON(500, gin.H{
			"details": "something went wrong, sorry",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":       wallet.GetId(),
		"user":     wallet.GetUserId(),
		"currency": wallet.GetCurrency(),
		"balance":  wallet.GetBalance(),
	})
}
