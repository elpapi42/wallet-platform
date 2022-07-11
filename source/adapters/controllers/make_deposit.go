package controllers

import (
	"wallet/source/adapters"
	"wallet/source/application"
	"wallet/source/infrastructure/eventstore"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MakeDepositSchema struct {
	Wallet uuid.UUID
	Amount float64
}

func MakeDepositController(c *gin.Context) {
	var payload MakeDepositSchema
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"details": err.Error(),
		})
		return
	}

	wallet_repo := &adapters.EventstoreWalletRepository{EventstoreClient: eventstore.GetClient()}

	service := application.MakeDepositService{WalletRepository: wallet_repo}

	err = service.Execute(payload.Wallet, payload.Amount)
	if err == esdb.ErrStreamNotFound {
		c.JSON(404, gin.H{
			"details": "wallet not found",
		})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{
			"details": "something went wrong, sorry",
		})
		return
	}
}
