package validations

import (
	"MongodbTest/model"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// BotCreateBody ...
func BotCreateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var botCreateBody model.BotCreateBody

		// bind request data
		if err := c.Bind(&botCreateBody); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := botCreateBody.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("botCreateBody", botCreateBody)

		return next(c)
	}
}

// BotUpdateBody ...
func BotUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var botUpdateBody model.BotUpdateBody

		// bind request data
		if err := c.Bind(&botUpdateBody); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := botUpdateBody.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("botUpdateBody", botUpdateBody)

		return next(c)
	}
}
