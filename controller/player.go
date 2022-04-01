package controller

import (
	"MongodbTest/config"
	"MongodbTest/model"
	"MongodbTest/service"
	"MongodbTest/util"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func MyProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	claims := &util.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().Jwt.SecretKey), nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(claims.Data["id"])

	// Process data
	doc, err := service.MyProfile(claims.Data["id"].(string))

	//// Process data
	//doc, err := service.MyProfile(claims["ID"].(string))

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"name":  doc.Name,
		"email": doc.Email,
	}

	// success
	return util.Response200(c, data, "")
}

// UpdateMyProfile ...
func UpdateMyProfile(c echo.Context) error {
	//// get player id
	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(jwt.MapClaims)
	//ID := claims["ID"].(string)

	user := c.Get("user").(*jwt.Token)

	claims := &util.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().Jwt.SecretKey), nil
	})
	if err != nil {
		panic(err)
	}

	ID := claims.Data["id"].(string)

	var playerUpdateBody = c.Get("playerRequestBody").(model.Player)

	// UpdateProfile
	err = service.UpdateProfile(ID, playerUpdateBody)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, ID, "")
}
