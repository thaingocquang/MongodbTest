package validations

import (
	"MongodbTest/model"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// PlayerRegisterBody ...
func PlayerRegisterBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var playerRegisterBody model.PlayerRegisterBody

		// bind request data
		if err := c.Bind(&playerRegisterBody); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := playerRegisterBody.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("playerRegisterBody", playerRegisterBody)

		return next(c)
	}
}

// PlayerLoginBody ...
func PlayerLoginBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var playerRegisterBody model.PlayerLoginBody

		// bind request data
		if err := c.Bind(&playerRegisterBody); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := playerRegisterBody.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("playerLoginBody", playerRegisterBody)

		return next(c)
	}
}

// ValidatePlayerUpdateBody ...
func ValidatePlayerUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var player model.Player

		// bind request body
		err := c.Bind(&player)

		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// success
		c.Set("playerRequestBody", player)

		return next(c)
	}
}
