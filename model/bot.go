package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
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
		MinBet       int                `bson:"minBet,omitempty"`
		MaxBet       int                `bson:"maxBet,omitempty"`
		CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}

	// BotCreateBody ...
	BotCreateBody struct {
		Name        string `json:"name"`
		TotalPoints int    `json:"totalPoints"`
		MinBet      int    `json:"minBet"`
		MaxBet      int    `json:"maxBet"`
	}

	// BotUpdateBody ...
	BotUpdateBody struct {
		Name         string `json:"name"`
		TotalPoints  int    `json:"totalPoints"`
		RemainPoints int    `json:"remainPoints"`
		MinBet       int    `json:"minBet"`
		MaxBet       int    `json:"maxBet"`
	}
)

// Validate ...
func (b BotCreateBody) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.TotalPoints, validation.Required),
		validation.Field(&b.MinBet, validation.Required),
		validation.Field(&b.MaxBet, validation.Required),
	)
}

// Validate ...
func (b BotUpdateBody) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Length(3, 30)),
		validation.Field(&b.TotalPoints),
		validation.Field(&b.RemainPoints),
		validation.Field(&b.MinBet),
		validation.Field(&b.MaxBet),
	)
}
