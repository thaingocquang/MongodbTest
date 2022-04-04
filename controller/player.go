package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
	"strconv"
)

// MyProfile ...
func MyProfile(c echo.Context) error {
	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)

	// player id
	playerID := jwtPayload["id"].(string)

	// get player profile
	profile, err := service.MyProfile(playerID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// get player stats
	stats, err := service.GetPlayerStatsByID(playerID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// response data
	data := map[string]interface{}{
		"profile": profile,
		"stats":   stats,
	}

	// success
	return util.Response200(c, data, "")
}

// UpdateMyProfile ...
func UpdateMyProfile(c echo.Context) error {
	var playerUpdateBody = c.Get("playerRequestBody").(model.Player)

	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)
	id := jwtPayload["id"].(string)

	// UpdateProfile
	err := service.UpdateProfile(id, playerUpdateBody)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, id, "")
}

// DeletePlayer ...
func DeletePlayer(c echo.Context) error {
	var strID = c.Get("strID").(string)

	// process
	if err := service.DeletePlayer(strID); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

// GetPlayerByID ...
func GetPlayerByID(c echo.Context) error {
	var strID = c.Get("strID").(string)
	nilPlayer := model.Player{}

	player := service.GetPlayerByID(strID)

	if player == nilPlayer {
		return util.Response404(c, nil, "")
	}

	return util.Response200(c, player, "")
}

// GetListPlayer ...
func GetListPlayer(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	players := service.GetListPlayer(page, limit)

	return util.Response200(c, players, "")
}
