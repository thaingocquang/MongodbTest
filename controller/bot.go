package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// CreateBot ...
func CreateBot(c echo.Context) error {
	var botBody = c.Get("botCreateBody").(model.BotCreateBody)

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

// UpdateBot ...
func UpdateBot(c echo.Context) error {
	var (
		strID         = c.Get("strID").(string)
		botUpdateBody = c.Get("botUpdateBody").(model.BotUpdateBody)
	)

	// process data
	if err := service.UpdateBotByID(strID, botUpdateBody); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

func GetListBot(c echo.Context) error {
	bots := service.GetListBot()

	return util.Response200(c, bots, "")
}
