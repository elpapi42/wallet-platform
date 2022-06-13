package gin

import (
	"wallet/source/adapters/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.POST("/users", controllers.CreateUserController)
}
