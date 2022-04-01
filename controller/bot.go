package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// CreateBot ...
func CreateBot(c echo.Context) error {
	var botBody = c.Get("bot").(model.Bot)

	// process data
	err := service.CreateBot(botBody)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

// GetBotByID ...
func GetBotByID(c echo.Context) error {
	var strID = c.Get("strID").(string)

	// process
	bot, err := service.GetBotByID(strID)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bot, "")
}
