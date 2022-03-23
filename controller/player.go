package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func MyProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Process data
	doc, err := service.MyProfile(claims["ID"].(string))

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"name":  doc.Name,
		"email": doc.Email,
	}

	// success
	return util.Response200(c, data, "")
}

// UpdateMyProfile ...
func UpdateMyProfile(c echo.Context) error {
	// get player id
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["ID"].(string)

	var playerUpdateBody = c.Get("playerRequestBody").(model.Player)

	// UpdateProfile
	err := service.UpdateProfile(ID, playerUpdateBody)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, ID, "")
}
