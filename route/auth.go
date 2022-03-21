package route

import (
	"MongodbTest/controller"
	"MongodbTest/validation"
	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) {
	players := e.Group("/auth")
	players.POST("/register", controller.Register, validation.RegisterBody)
	players.POST("/login", nil)
}
