package validations

import (
	"MongodbTest/model"
	"errors"
	"github.com/labstack/echo/v4"
)

func ValidateGameValue(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var gameBody model.GameBody

		// bind
		err := c.Bind(&gameBody)

		if err != nil {
			return errors.New("bind request failed")
		}

		c.Set("game", gameBody)

		return next(c)
	}
}
