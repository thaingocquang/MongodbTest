package middleware

import (
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
