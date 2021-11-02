package sql_store

import "github.com/reqww/go-rest-api/internal/app/auth"

type AuthDataRepository struct {
	store *Store
}

func (a *AuthDataRepository) Create(authData *auth.UserAuthData) error {
	return nil
}

func (a *AuthDataRepository) All() ([]*auth.UserAuthData, error) {
	return nil, nil
}