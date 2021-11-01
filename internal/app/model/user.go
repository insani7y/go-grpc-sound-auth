package model

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/reqww/go-rest-api/internal/app/fileHandler"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}


func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(RequiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}


func (u *User) GetAllThingsDone(filesBytes [][]byte) error {

	handler := fileHandler.New()

	for _, fileBytes := range filesBytes {
		data, err := handler.DetermineAmplitudeValues(fileBytes)
		if err != nil {
			return err
		}

		frames := handler.FrameCut(data)

		amplitudes := handler.FourierTransform(frames)

		melAmplitudes := handler.ToMelScale(amplitudes)

		features := handler.GetMelFeaturesArr(melAmplitudes)

		fmt.Println(features)
		fmt.Println(len(features))
	}



	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}
