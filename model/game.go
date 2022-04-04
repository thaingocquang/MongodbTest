package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Game ...
	Game struct {
		ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		GameNo     int                `json:"gameNo,omitempty" bson:"gameNo,omitempty"`
		PlayerID   primitive.ObjectID `json:"playerID,omitempty" bson:"playerID,omitempty"`
		BotID      primitive.ObjectID `json:"botID,omitempty" bson:"botID,omitempty"`
		WinnerID   primitive.ObjectID `json:"winnerID,omitempty" bson:"winnerID,omitempty"`
		PlayerHand Hand               `json:"playerHand,omitempty" bson:"playerHand,omitempty"`
		BotHand    Hand               `json:"botHand,omitempty" bson:"botHand,omitempty"`
		BetValue   int                `json:"betValue,omitempty" bson:"betValue,omitempty"`
		CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}

	// GameBody ...
	GameBody struct {
		BetValue int `json:"betValue"`
	}
)
