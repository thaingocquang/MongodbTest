package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBot ...
func CreateBot(bot model.Bot) error {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)

	// InsertOne
	_, err := botCol.InsertOne(ctx, bot)

	return err
}

// GetBotByID ...
func GetBotByID(botID primitive.ObjectID) (model.Bot, error) {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
		bot    model.Bot
	)

	filter := bson.M{"_id": botID}

	// FindOne
	if err := botCol.FindOne(ctx, filter).Decode(&bot); err != nil {
		return bot, err
	}

	return bot, nil
}

// GetBotByName ...
func GetBotByName(botName string) (model.Bot, error) {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
		bot    model.Bot
	)

	filter := bson.M{"name": botName}

	// FindOne
	err := botCol.FindOne(ctx, filter).Decode(&bot)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

// UpdateBotByID ...
func UpdateBotByID(id primitive.ObjectID, bot model.Bot) error {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)

	// UpdateOne
	_, err := botCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bot})
	if err != nil {
		return err
	}

	return nil
}

// UpdateBotByName ...
func UpdateBotByName(name string, bot model.Bot) error {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)

	// UpdateOne
	_, err := botCol.UpdateOne(ctx, bson.M{"name": name}, bson.M{"$set": bot})
	if err != nil {
		return err
	}

	return nil
}

func GetListBot() []model.Bot {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
		bots   []model.Bot
	)

	cursor, err := botCol.Find(ctx, bson.D{})
	if err != nil {
		return nil
	}

	if err = cursor.All(ctx, &bots); err != nil {
		return nil
	}

	return bots
}
