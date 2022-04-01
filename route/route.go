package route

import "github.com/labstack/echo/v4"

// Route ...
func Route(e *echo.Echo) {
	auth(e)
	player(e)
	game(e)
	bot(e)
}
