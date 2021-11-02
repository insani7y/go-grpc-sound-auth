package store

import (
	"github.com/reqww/go-rest-api/internal/app/auth"
	"github.com/reqww/go-rest-api/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
	All() ([]*model.User, error)
}

type AuthDataRepository interface {
	Create(*auth.UserAuthData) error
	All() ([]*auth.UserAuthData, error)
}
