package controller

import (
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// Register ...
func Register(c echo.Context) error {
	var registerBody = c.Get("registerBody").(model.PlayerRegister)

	// process data
	err := service.Register(registerBody)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}
