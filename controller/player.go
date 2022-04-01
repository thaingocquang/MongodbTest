package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// MyProfile ...
func MyProfile(c echo.Context) error {
	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)

	// Process data
	doc, err := service.MyProfile(jwtPayload["id"].(string))

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// response data
	data := map[string]interface{}{
		"name":  doc.Name,
		"email": doc.Email,
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
