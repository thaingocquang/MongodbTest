package middleware

import (
	"MongodbTest/model"
	"MongodbTest/util"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
)

// ValidatePlayerRegisterBody ...
func ValidatePlayerRegisterBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var player model.Player

		// bind request body
		c.Bind(&player)

		// validate
		err := validation.Errors{
			"name":     validation.Validate(player.Name, validation.Required),
			"email":    validation.Validate(player.Email, validation.Required, is.Email),
			"password": validation.Validate(player.Password, validation.Required),
		}.Filter()

		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// success
		c.Set("playerRequestBody", player)

		return next(c)
	}
}

// ValidatePlayerLoginBody ...
func ValidatePlayerLoginBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var player model.Player

		// bind request body
		c.Bind(&player)

		// validate
		err := validation.Errors{
			"email":    validation.Validate(player.Email, validation.Required),
			"password": validation.Validate(player.Password, validation.Required),
		}.Filter()

		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// success
		c.Set("playerRequestBody", player)

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
