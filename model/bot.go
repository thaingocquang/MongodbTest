package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Bot ...
	Bot struct {
		ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Name         string             `json:"name,omitempty" bson:"name,omitempty"`
		TotalPoints  int                `json:"totalPoints,omitempty" bson:"totalPoints,omitempty"`
		RemainPoints int                `json:"remainPoints,omitempty" bson:"remainPoints,omitempty"`
		CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}
)
