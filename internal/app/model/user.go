package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	UserId int `json:"id"`
	Email string `json:"email"`
}


func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

