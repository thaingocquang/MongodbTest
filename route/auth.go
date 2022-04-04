package route

import (
	"MongodbTest/controller"
	"MongodbTest/validations"
	"github.com/labstack/echo/v4"
)

// auth ...
func auth(e *echo.Echo) {
	players := e.Group("/auth")
	players.POST("/register", controller.Register, validations.PlayerRegisterBody)
	players.POST("/login", controller.Login, validations.PlayerLoginBody)
	players.POST("/admin-login", controller.AdminLogin, validations.AdminLoginBody)
}
