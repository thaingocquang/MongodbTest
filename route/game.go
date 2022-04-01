package route

import (
	"MongodbTest/controller"
	customMiddleware "MongodbTest/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// game ...
func game(e *echo.Echo) {
	game := e.Group("/games", middleware.JWT([]byte(envVars.Jwt.SecretKey)))
	game.POST("", controller.Play, customMiddleware.ValidateGameValue)
}
