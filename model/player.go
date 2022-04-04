package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Player ...
	Player struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Name      string             `bson:"name,omitempty"`
		Email     string             `bson:"email,omitempty"`
		Password  string             `bson:"password,omitempty"`
		CreatedAt time.Time          `bson:"created_at,omitempty"`
		UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	}

	// PlayerRegisterBody ...
	PlayerRegisterBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// PlayerLoginBody ...
	PlayerLoginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (p PlayerRegisterBody) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.Password, validation.Required),
	)
}

func (p PlayerLoginBody) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.Password, validation.Required),
	)
}
