package middleware

import (
	"MongodbTest/config"
	"MongodbTest/util"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)

		claims := &util.JwtCustomClaims{}
		_, err := jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv().Jwt.SecretKey), nil
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(claims.Data["id"])
		fmt.Println(claims.Data["isAdmin"])

		if claims.Data["isAdmin"] == true {
			return next(c)
		}

		return util.Response400(c, nil, "authorization fail: not admin")

	}
}
