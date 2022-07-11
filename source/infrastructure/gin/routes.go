package gin

import (
	"wallet/source/adapters/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.POST("/wallets", controllers.CreateWalletController)
	engine.POST("/deposits", controllers.MakeDepositController)
}
