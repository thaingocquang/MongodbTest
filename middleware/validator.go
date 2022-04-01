package middleware

import (
	"MongodbTest/model"
	"MongodbTest/util"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"unicode"
)

// ValidatePlayerRegisterBody ...
func ValidatePlayerRegisterBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var player model.Player

		// bind request body
		c.Bind(&player)

		// validate
		//err := validation.Errors{
		//	"name":     validation.Validate(player.Name, validation.Required),
		//	"email":    validation.Validate(player.Email, validation.Required, is.Email),
		//	"password": validation.Validate(player.Password, validation.Required),
		//}.Filter()

		var (
			hasMinLen  = false
			hasUpper   = false
			hasLower   = false
			hasNumber  = false
			hasSpecial = false
			err        error
		)

		err = validation.ValidateStruct(&player,
			validation.Field(&player.Name, validation.Required),
			validation.Field(&player.Email, validation.Required, is.Email),
			validation.Field(&player.Password, validation.Required),
		)

		if len(player.Password) >= 7 {
			hasMinLen = true
		}
		for _, char := range player.Password {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
			}
		}

		if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
			err = errors.New("password invalid")
		}

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

//
func ValidateCreateBotBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var bot model.Bot

		// bind
		err := c.Bind(&bot)

		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("bot", bot)

		return next(c)
	}
}

func ValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		_, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("strID", id)

		return next(c)
	}
}
