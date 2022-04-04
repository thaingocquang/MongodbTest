package validations

import (
	"MongodbTest/model"
	"MongodbTest/util"
	"github.com/labstack/echo/v4"
)

// CheckAdminRole ...
func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtPayload, _ := util.GetJWTPayload(c)

		if jwtPayload["isAdmin"] == true {
			return next(c)
		}

		return util.Response400(c, nil, "authorization fail: not admin")
	}
}

// AdminLoginBody ...
func AdminLoginBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var adminLoginBody model.AdminLoginBody

		// bind request data
		if err := c.Bind(&adminLoginBody); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := adminLoginBody.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("adminLoginBody", adminLoginBody)

		return next(c)
	}
}
