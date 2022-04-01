package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBot(bot model.Bot) error {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)

	// InsertOne
	_, err := botCol.InsertOne(ctx, bot)

	return err
}

func GetBotByID(botID primitive.ObjectID) (model.Bot, error) {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
		bot    model.Bot
	)

	filter := bson.M{"_id": botID}

	err := botCol.FindOne(ctx, filter).Decode(&bot)

	if err != nil {
		return bot, err
	}

	return bot, nil

}

func GetBotByName(botName string) (model.Bot, error) {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
		bot    model.Bot
	)

	filter := bson.M{"name": botName}

	err := botCol.FindOne(ctx, filter).Decode(&bot)

	if err != nil {
		return bot, err
	}

	return bot, nil

}

func UpdateBotByID(id string, bot model.Bot) error {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := botCol.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bot})
	if err != nil {
		return err
	}

	return nil
}

func UpdateBotByName(name string, bot model.Bot) error {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)

	_, err := botCol.UpdateOne(ctx, bson.M{"name": name}, bson.M{"$set": bot})
	if err != nil {
		return err
	}

	return nil
}
