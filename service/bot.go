package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateBot(bot model.Bot) error {
	// check exist name

	bot.ID = primitive.NewObjectID()
	bot.CreatedAt = time.Now()
	bot.UpdatedAt = time.Now()

	// create bot
	err := dao.CreateBot(bot)

	if err != nil {
		return errors.New("can not create bot")
	}

	return err
}

func GetBotByID(botID string) (model.Bot, error) {
	objID, _ := primitive.ObjectIDFromHex(botID)

	bot, err := dao.GetBotByID(objID)

	if err != nil {
		return bot, err
	}

	return bot, nil
}
