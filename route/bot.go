package route

import (
	"MongodbTest/controller"
	"MongodbTest/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// bot ...
func bot(e *echo.Echo) {
	players := e.Group("/bots", middleware.JWT([]byte(envVars.Jwt.SecretKey)), validations.CheckAdminRole)

	players.POST("", controller.CreateBot, validations.BotCreateBody)
	players.GET("/:id", controller.GetBotByID, validations.ValidateID)
	players.GET("", controller.GetListBot, validations.CheckAdminRole)
	players.PUT("/:id", controller.UpdateBot, validations.ValidateID, validations.BotUpdateBody)
}
