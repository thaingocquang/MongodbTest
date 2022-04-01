package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// Play ...
func Play(c echo.Context) error {
	// get game info (bot name, bet value)
	var gameBody = c.Get("game").(model.GameBody)

	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)

	// process data
	err := service.Play(gameBody, jwtPayload["id"].(string))
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}
