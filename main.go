package main

import (
	"MongodbTest/config"
	"MongodbTest/dao"
	"MongodbTest/module/database"
	"MongodbTest/route"
	"github.com/labstack/echo/v4"
)

// init ...
func init() {
	config.Init()
	database.Connect()
	dao.InitAdminUser()
}

func main() {
	// envVars ...
	envVars := config.GetEnv()

	// echo ...
	e := echo.New()

	// route
	route.Route(e)

	// start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}
