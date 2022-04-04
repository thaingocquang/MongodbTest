package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// Play ...
func Play(c echo.Context) error {
	// get game value from context (bot id, bet value)
	var gameBody = c.Get("game").(model.GameBody)
	var botID = c.Get("strID").(string)

	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)

	// process data
	game, err := service.Play(gameBody, jwtPayload["id"].(string), botID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, game, "")
}
