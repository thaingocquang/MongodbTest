package route

import (
	"MongodbTest/controller"
	"MongodbTest/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// game ...
func game(e *echo.Echo) {
	game := e.Group("/games", middleware.JWT([]byte(envVars.Jwt.SecretKey)))
	game.POST("/:id", controller.Play, validations.ValidateID, validations.ValidateGameValue)
}
