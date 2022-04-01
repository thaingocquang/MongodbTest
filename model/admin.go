package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Admin struct {
		ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Username string             `json:"username,omitempty" bson:"username,omitempty"`
		Password string             `json:"password,omitempty" bson:"password,omitempty"`
	}
)
