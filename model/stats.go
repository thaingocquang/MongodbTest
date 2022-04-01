package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Stats ...
	Stats struct {
		ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		PlayerID  primitive.ObjectID `json:"playerID,omitempty" bson:"playerID,omitempty"`
		Point     int                `json:"point,omitempty" bson:"point,omitempty"`
		TotalGame int                `json:"totalGame,omitempty" bson:"totalGame,omitempty"`
		WinGame   int                `json:"winGame,omitempty" bson:"winGame,omitempty"`
		WinRate   float32            `json:"winRate,omitempty" bson:"winRate,omitempty"`
		CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}
)
