package route

import (
	"MongodbTest/controller"
	"MongodbTest/middleware"
	"github.com/labstack/echo/v4"
)

// auth ...
func auth(e *echo.Echo) {
	players := e.Group("/auth")
	players.POST("/register", controller.Register, middleware.ValidatePlayerRegisterBody)
	players.POST("/login", controller.Login, middleware.ValidatePlayerLoginBody)
	players.POST("/admin-login", controller.AdminLogin)
}
