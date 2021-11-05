package store

import (
	"github.com/reqww/go-rest-api/internal/app/auth"
	"github.com/reqww/go-rest-api/internal/app/model"
	"mime/multipart"
)

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
}

type AuthDataRepository interface {
	Create(*auth.UserAuthData) error
	All() ([]*auth.UserAuthData, error)
	SaveMFCC([]multipart.File, string, int)
	CreateAuthData(multipart.File, string, int)
	DetermineUserBySound([][]float64) (int, error)
}
