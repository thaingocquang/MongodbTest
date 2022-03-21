package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"time"
)

type (
	// PlayerRegister ...
	PlayerRegister struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// PlayerRegisterBSON ...
	PlayerRegisterBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		Name           string             `bson:"name"`
		Email          string             `bson:"email"`
		HashedPassword string             `bson:"hashedPassword"`
		CreatedAt      time.Time          `bson:"createdAt"`
		UpdatedAt      time.Time          `bson:"updatedAt"`
	}

	// PlayerLogin ...
	PlayerLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (body PlayerRegister) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.Email,
			validation.Required,
			validation.Match(regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)),
		),
		validation.Field(
			&body.Password,
			validation.Required,
		),
	)

	//if len(body.Name) == 0 {
	//	return errors.New("name required")
	//}
	//
	//if len(body.Email) == 0 {
	//	return errors.New("email required")
	//}
	//
	//if len(body.Password) == 0 {
	//	return errors.New("password required")
	//}
	//
	//if !strings.ContainsAny(body.Email, "@") {
	//	return errors.New("at sign @ required")
	//}
	//
	//return nil
}

func (body PlayerLogin) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.Email,
			validation.Required,
			validation.Match(regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)),
		),
		validation.Field(
			&body.Password,
			validation.Required,
		),
	)
}
