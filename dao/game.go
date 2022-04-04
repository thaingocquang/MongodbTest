package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// RecordGame ...
func RecordGame(game model.Game) error {
	var (
		gameCol = database.GameCol()
		ctx     = context.Background()
	)

	// InsertOne
	_, err := gameCol.InsertOne(ctx, game)

	return err
}

// CountAllGame ...
func CountAllGame() int {
	var (
		statsCol = database.StatsCol()
		ctx      = context.Background()
	)
	count, err := statsCol.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}
	return int(count)
}
