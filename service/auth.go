package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
	"MongodbTest/util"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Register ...
func Register(registerBody model.Player) error {
	// check email existed
	isEmailExisted := dao.CheckPlayerEmailExisted(registerBody.Email)
	if isEmailExisted {
		err := errors.New("email already existed")
		return err
	}

	// hash player password
	bytes, err := bcrypt.GenerateFromPassword([]byte(registerBody.Password), 14)
	if err != nil {
		return err
	}

	registerBody.ID = primitive.NewObjectID()
	registerBody.Password = string(bytes)
	registerBody.CreatedAt = time.Now()
	registerBody.UpdatedAt = time.Now()

	// create player
	err = dao.PlayerCreate(registerBody)
	if err != nil {
		err := errors.New("can't create player")
		return err
	}

	return err
}

// Login ...
func Login(loginBody model.Player) (string, error) {
	// find player in db
	player, err := dao.PlayerFindByEmail(loginBody.Email)

	if err != nil {
		return "", err
	}

	// verify player password
	if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(loginBody.Password)); err != nil {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id": player.ID,
	}

	// return JWT token
	return util.GenerateUserToken(data), err
}
