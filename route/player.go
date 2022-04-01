package route

import (
	"MongodbTest/config"
	"MongodbTest/controller"
	customMiddleware "MongodbTest/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVars = config.GetEnv()

// player ...
func player(e *echo.Echo) {
	players := e.Group("/players", middleware.JWT([]byte(envVars.Jwt.SecretKey)))
	players.GET("/me", controller.MyProfile)
	players.PUT("/me", controller.UpdateMyProfile, customMiddleware.ValidatePlayerUpdateBody)
	players.DELETE("", nil)
	players.GET("/id", nil)
	players.GET("", nil, customMiddleware.CheckAdminRole)
}
