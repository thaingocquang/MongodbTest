package route

import (
	"MongodbTest/controller"
	"MongodbTest/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// bot ...
func bot(e *echo.Echo) {
	players := e.Group("/bots", middleware.JWT([]byte(envVars.Jwt.SecretKey)))

	players.POST("", controller.CreateBot, validations.CheckAdminRole, validations.BotCreateBody)
	players.GET("/:id", controller.GetBotByID, validations.ValidateID)
	players.GET("", controller.GetListBot)
	players.PUT("/:id", controller.UpdateBot, validations.CheckAdminRole, validations.ValidateID, validations.BotUpdateBody)
}
