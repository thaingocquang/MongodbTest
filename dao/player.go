package dao

import (
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

// PlayerFindByEmail ...
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

// PlayerProfileFindByID ...
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

// UpdatePlayerProfile ...
func UpdatePlayerProfile(ID string, newProfile model.Player) error {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
		//profile   model.Player
	)

	objID, _ := primitive.ObjectIDFromHex(ID)
	update := model.Player{Name: newProfile.Name, Email: newProfile.Email, Password: newProfile.Password}

	// UpdateOne
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

// DeletePlayer ...
func DeletePlayer(ID primitive.ObjectID) error {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)

	if _, err := playerCol.DeleteOne(ctx, bson.D{{"_id", ID}}); err != nil {
		return err
	}

	return nil
}

func GetPlayerByID(ID primitive.ObjectID) model.Player {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
		player    model.Player
	)

	filter := bson.M{"_id": ID}

	// FindOne
	if err := playerCol.FindOne(ctx, filter).Decode(&player); err != nil {
		return model.Player{}
	}

	return player
}

func GetListPlayer(page, limit int) []model.Player {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
		players   []model.Player
	)

	options := new(options2.FindOptions)
	if limit != 0 {
		if page == 0 {
			page = 1
		}
		options.SetSkip(int64((page - 1) * limit))
		options.SetLimit(int64(limit))
	}

	cursor, err := playerCol.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil
	}

	if err = cursor.All(ctx, &players); err != nil {
		return nil
	}

	return players
}
