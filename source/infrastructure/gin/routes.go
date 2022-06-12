package gin

import (
	"wallet/source/adapters"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.POST("/users", adapters.CreateUserController)
}
