package test_helper

import (
	"MongodbTest/config"
	"MongodbTest/route"
	"github.com/labstack/echo/v4"
)

// InitServer ...
func InitServer() *echo.Echo {
	// Init initialize app's config
	config.Init()

	// ConnectTestDB ...
	ConnectTestDB()

	// ClearDB ...
	ClearDB()

	// new test server
	e := echo.New()

	// route
	route.Route(e)

	return e
}
