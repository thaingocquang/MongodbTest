package route

import (
	"MongodbTest/controller"
	customMiddleware "MongodbTest/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// bot ...
func bot(e *echo.Echo) {
	players := e.Group("/bot", middleware.JWT([]byte(envVars.Jwt.SecretKey)), customMiddleware.CheckAdminRole)
	players.POST("/", controller.CreateBot, customMiddleware.ValidateCreateBotBody)
	players.GET("/:id", controller.GetBotByID, customMiddleware.ValidateID)
	players.PUT("/", controller.AdminLogin)
}
