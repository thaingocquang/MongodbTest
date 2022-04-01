package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func AdminFindByUsername(username string) (model.Admin, error) {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
		admin    model.Admin
	)

	// find player
	filter := bson.M{"username": username}
	err := adminCol.FindOne(ctx, filter).Decode(&admin)

	// if err
	if err != nil {
		return admin, err
	}

	return admin, nil
}
