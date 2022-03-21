package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckPlayerEmailExisted(email string) bool {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)

	// count documents having same email
	count, err := playerCol.CountDocuments(ctx, bson.M{"email": email})
	if err != nil || count <= 0 {
		return false
	}

	return true
}

func PlayerCreate(doc model.PlayerRegisterBSON) error {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)

	// InsertOne
	_, err := playerCol.InsertOne(ctx, doc)

	return err
}
