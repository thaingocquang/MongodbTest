package route

import "github.com/labstack/echo/v4"

// player ...
func player(e *echo.Echo) {
	players := e.Group("/players")
	players.GET("/profile", nil)
}
