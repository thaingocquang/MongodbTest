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
func Register(registerBody model.PlayerRegisterBody) error {
	// check email existed
	isEmailExisted := dao.CheckPlayerEmailExisted(registerBody.Email)
	if isEmailExisted {
		return errors.New("email already existed")
	}

	// hash player password
	bytes, err := bcrypt.GenerateFromPassword([]byte(registerBody.Password), 14)
	if err != nil {
		return err
	}

	// player objectID
	objID := primitive.NewObjectID()

	// playerRegisterBSON
	playerRegisterBSON := model.Player{
		ID:        objID,
		Name:      registerBody.Name,
		Email:     registerBody.Email,
		Password:  string(bytes),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// create player
	if err := dao.PlayerCreate(playerRegisterBSON); err != nil {
		return errors.New("can't create player")
	}

	// Create empty player stats
	playerStats := model.StatsCreate{
		ID:        primitive.NewObjectID(),
		PlayerID:  objID,
		Point:     0,
		TotalGame: 0,
		WinGame:   0,
		WinRate:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := dao.CreatePlayerStats(playerStats); err != nil {
		return errors.New("can't create player stats")
	}

	return err
}

// Login ...
func Login(playerLoginBody model.PlayerLoginBody) (string, error) {
	// find player in db
	player, err := dao.PlayerFindByEmail(playerLoginBody.Email)
	if err != nil {
		return "", err
	}

	// verify player password
	if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(playerLoginBody.Password)); err != nil {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id": player.ID,
	}

	// return JWT token
	return util.GenerateUserToken(data), err
}

// AdminLogin ...
func AdminLogin(loginBody model.AdminLoginBody) (string, error) {
	// find admin in db
	admin, err := dao.AdminFindByUsername(loginBody.Username)

	if err != nil {
		return "", err
	}

	// verify admin password
	if admin.Password != loginBody.Password {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id":      admin.ID,
		"isAdmin": true,
	}

	// return JWT token
	return util.GenerateUserToken(data), err
}
