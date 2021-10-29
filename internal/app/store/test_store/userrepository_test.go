package test_store

import (
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := New()
	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}


func TestUserRepository_FindByEmail(t *testing.T) {
	s := New()
	email := "sas@sas.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	err = s.User().Create(u)
	assert.NoError(t, err)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindById(t *testing.T) {
	s := New()
	u := model.TestUser(t)

	err := s.User().Create(u)
	assert.NoError(t, err)

	u, err = s.User().FindById(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

