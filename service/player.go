package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
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
