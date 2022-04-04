package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateStatsByPlayerID ...
func UpdateStatsByPlayerID(id primitive.ObjectID, stats model.Stats) error {
	var (
		statsCol = database.StatsCol()
		ctx      = context.Background()
	)

	// UpdateOne
	u, err := statsCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": stats})
	if err != nil {
		return err
	}

	fmt.Println("UPDATED: ", u.MatchedCount)

	return nil
}

// GetStatsByPlayerID ...
func GetStatsByPlayerID(id primitive.ObjectID) (model.Stats, error) {
	var (
		statsCol = database.StatsCol()
		ctx      = context.Background()
		stats    model.Stats
	)

	filter := bson.M{"playerID": id}
	err := statsCol.FindOne(ctx, filter).Decode(&stats)

	// if err
	if err != nil {
		return stats, err
	}

	return stats, nil

}

// CreatePlayerStats ...
func CreatePlayerStats(stats model.StatsCreate) error {
	var (
		statsCol = database.StatsCol()
		ctx      = context.Background()
	)

	// InsertOne
	_, err := statsCol.InsertOne(ctx, stats)

	return err
}
