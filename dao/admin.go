package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AdminFindByUsername ...
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

func InitAdminUser() {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
	)

	count, _ := adminCol.CountDocuments(ctx, bson.D{})

	if count == 0 {
		admin := model.Admin{
			ID:       primitive.NewObjectID(),
			Username: "admin",
			Password: "123456",
		}
		adminCol.InsertOne(ctx, admin)
	}
}
