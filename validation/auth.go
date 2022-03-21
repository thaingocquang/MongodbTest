package validation

import (
	"MongodbTest/model"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

func RegisterBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var registerBody model.PlayerRegister

		// bind request body
		c.Bind(&registerBody)

		// validate
		if err := registerBody.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// success
		c.Set("registerBody", registerBody)

		return next(c)
	}
}
