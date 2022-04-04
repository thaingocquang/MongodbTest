package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// CreateBot ...
func CreateBot(botCreateBody model.BotCreateBody) error {
	// bot BSON
	bot := model.Bot{
		ID:           primitive.NewObjectID(),
		Name:         botCreateBody.Name,
		TotalPoints:  botCreateBody.TotalPoints,
		RemainPoints: botCreateBody.TotalPoints,
		MinBet:       botCreateBody.MinBet,
		MaxBet:       botCreateBody.MaxBet,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// create bot
	if err := dao.CreateBot(bot); err != nil {
		return errors.New("can not create bot")
	}

	return nil
}

// GetBotByID ...
func GetBotByID(botID string) (model.Bot, error) {
	// to objectID
	objID, _ := primitive.ObjectIDFromHex(botID)

	// get bot by id
	bot, err := dao.GetBotByID(objID)

	if err != nil {
		return bot, err
	}

	return bot, nil
}

// UpdateBotByID ...
func UpdateBotByID(ID string, botUpdateBody model.BotUpdateBody) error {
	// bot BSON
	bot := model.Bot{
		Name:         botUpdateBody.Name,
		TotalPoints:  botUpdateBody.TotalPoints,
		RemainPoints: botUpdateBody.RemainPoints,
		MinBet:       botUpdateBody.MinBet,
		MaxBet:       botUpdateBody.MaxBet,
		UpdatedAt:    time.Now(),
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	return dao.UpdateBotByID(objID, bot)
}

// GetListBot ...
func GetListBot() []model.Bot {
	return dao.GetListBot()
}
