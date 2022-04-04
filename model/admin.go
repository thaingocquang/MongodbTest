package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// Admin ...
	Admin struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Username string             `bson:"username,omitempty"`
		Password string             `bson:"password,omitempty"`
	}

	// AdminLoginBody ...
	AdminLoginBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}
)

func (a AdminLoginBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}
