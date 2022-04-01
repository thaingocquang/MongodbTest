package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateStatsByPlayerID(id string, stats model.Stats) error {
	var (
		statsCol = database.StatsCol()
		ctx      = context.Background()
		//profile   model.Player
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := statsCol.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": stats})
	if err != nil {
		return err
	}

	return nil
}

func GetStatsByPlayerID(id string) (model.Stats, error) {
	var (
		statsCol = database.StatsCol()
		ctx      = context.Background()
		stats    model.Stats
	)

	// objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"playerID": objID}
	err := statsCol.FindOne(ctx, filter).Decode(&stats)

	// if err
	if err != nil {
		return stats, err
	}

	return stats, nil

}
