package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PlayerFindByEmail(email string) (model.Player, error) {
	var (
		playerCol    = database.PlayerCol()
		ctx          = context.Background()
		playerDetail model.Player
	)

	// find player
	filter := bson.M{"email": email}
	err := playerCol.FindOne(ctx, filter).Decode(&playerDetail)

	// if err
	if err != nil {
		return playerDetail, err
	}

	return playerDetail, nil
}

func PlayerProfileFindByID(ID string) (model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
		profile   model.Player
	)

	// objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// find profile
	filter := bson.M{"_id": objID}
	err := playerCol.FindOne(ctx, filter).Decode(&profile)

	// if err
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func UpdatePlayerProfile(ID string, newProfile model.Player) error {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
		//profile   model.Player
	)

	objID, _ := primitive.ObjectIDFromHex(ID)

	//update := bson.M{"name": profile.Name, "email": profile.Email}

	update := model.Player{Name: newProfile.Name, Email: newProfile.Email, Password: newProfile.Password}

	//update := bson.M{"name": newProfile.Name, "email": newProfile.Email, "password": newProfile.Password}

	//res, err := playerCol.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

	_, err := playerCol.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

	if err != nil {
		return err
	}

	return nil
}

// CheckPlayerEmailExisted ...
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

// PlayerCreate ...
func PlayerCreate(doc model.Player) error {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)

	// InsertOne
	_, err := playerCol.InsertOne(ctx, doc)

	return err
}
