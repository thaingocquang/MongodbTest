package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MyProfile ...
func MyProfile(ID string) (model.Player, error) {
	doc, err := dao.PlayerProfileFindByID(ID)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

// UpdateProfile ...
func UpdateProfile(ID string, newProfile model.Player) error {
	err := dao.UpdatePlayerProfile(ID, newProfile)
	return err
}

// DeletePlayer ...
func DeletePlayer(ID string) error {
	objID, _ := primitive.ObjectIDFromHex(ID)
	return dao.DeletePlayer(objID)
}

// GetPlayerByID ...
func GetPlayerByID(ID string) model.Player {
	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	return dao.GetPlayerByID(objID)
}

// GetListPlayer ...
func GetListPlayer(page, limit int) []model.Player {
	return dao.GetListPlayer(page, limit)
}

// GetPlayerStatsByID ...
func GetPlayerStatsByID(id string) (model.Stats, error) {
	// to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	stats, err := dao.GetStatsByPlayerID(objID)
	if err != nil {
		return stats, err
	}

	return stats, nil
}
