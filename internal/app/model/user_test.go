package model_test

import (
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name string
		u func() *model.User
		isValid bool
	}{
		{name: "Valid", u: func() *model.User { return model.TestUser(t) }, isValid: true},
		{name: "EmptyEmail", u: func() *model.User {
			u := model.TestUser(t)
			u.Email = ""
			return u
		}, isValid: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			}
		})
	}
}