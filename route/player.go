package route

import (
	"MongodbTest/config"
	"MongodbTest/controller"
	"MongodbTest/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVars = config.GetEnv()

// player ...
func player(e *echo.Echo) {
	players := e.Group("/players", middleware.JWT([]byte(envVars.Jwt.SecretKey)))
	players.GET("/me", controller.MyProfile)
	players.PUT("/me", controller.UpdateMyProfile, validations.ValidatePlayerUpdateBody)

	players.DELETE("/:id", controller.DeletePlayer, validations.CheckAdminRole, validations.ValidateID)
	players.GET("/:id", controller.GetPlayerByID, validations.CheckAdminRole, validations.ValidateID)
	players.GET("", controller.GetListPlayer, validations.CheckAdminRole)
}
