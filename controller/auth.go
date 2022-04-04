package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

//type Authenticate interface {
//	ReturnToken(string, ) string
//}

//// Register ...
//func (a *Authenticate) Register(c echo.Context) error {
//	var registerBody = c.Get("playerRequestBody").(model.Player)
//
//	// process data
//	err := service.Register(registerBody)
//
//	// if err
//	if err != nil {
//		return util.Response400(c, nil, err.Error())
//	}
//
//	// success
//	return util.Response200(c, nil, "")
//}

// Register ...
func Register(c echo.Context) error {
	var playerRegisterBody = c.Get("playerRegisterBody").(model.PlayerRegisterBody)

	// process data
	err := service.Register(playerRegisterBody)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}

// Login ...
func Login(c echo.Context) error {
	var playerLoginBody = c.Get("playerLoginBody").(model.PlayerLoginBody)

	// process data
	token, err := service.Login(playerLoginBody)

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

// AdminLogin ...
func AdminLogin(c echo.Context) error {
	var admin = c.Get("adminLoginBody").(model.AdminLoginBody)

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
}
