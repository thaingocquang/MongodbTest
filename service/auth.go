package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(registerBody model.PlayerRegister) error {
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

	playerRegisterBSON := model.PlayerRegisterBSON{
		ID:             primitive.NewObjectID(),
		Name:           registerBody.Name,
		Email:          registerBody.Email,
		HashedPassword: string(bytes),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}

	// create player
	err = dao.PlayerCreate(playerRegisterBSON)
	if err != nil {
		err := errors.New("can't create player")
		return err
	}

	return err
}
