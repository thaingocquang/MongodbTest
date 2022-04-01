package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Player ...
	Player struct {
		ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Name      string             `json:"name,omitempty" bson:"name,omitempty"`
		Email     string             `json:"email,omitempty" bson:"email,omitempty"`
		Password  string             `json:"password,omitempty" bson:"password,omitempty"`
		CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}
)
