package util

import (
	"MongodbTest/config"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// jwtCustomClaims ...
type jwtCustomClaims struct {
	ID string
	jwt.StandardClaims
}

var envVars = config.GetEnv()

// GenerateUserToken ...
func GenerateUserToken(data map[string]interface{}) string {
	// claims ...
	claims := &jwtCustomClaims{
		data["id"].(primitive.ObjectID).Hex(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	st, err := token.SignedString([]byte(envVars.Jwt.SecretKey))

	// if err
	if err != nil {
		return ""
	}

	return st
}
