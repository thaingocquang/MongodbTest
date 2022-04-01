package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
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
