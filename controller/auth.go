package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// Register ...
func Register(c echo.Context) error {
	var registerBody = c.Get("playerRequestBody").(model.Player)

	// process data
	err := service.Register(registerBody)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}

// Login ...
func Login(c echo.Context) error {
	var loginBody = c.Get("playerRequestBody").(model.Player)

	// process data
	token, err := service.Login(loginBody)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token": token,
	}

	return util.Response200(c, data, "")
}

func AdminLogin(c echo.Context) error {
	var admin model.Admin

	// bind request body
	c.Bind(&admin)

	// process data
	token, err := service.AdminLogin(admin)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token":   token,
		"isAdmin": true,
	}

	return util.Response200(c, data, "")

	return nil
}
