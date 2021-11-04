package sql_store

import (
	"github.com/lib/pq"
	"github.com/reqww/go-rest-api/internal/app/auth"
	"mime/multipart"
)

type AuthDataRepository struct {
	store *Store
}

func (a *AuthDataRepository) Create(authData *auth.UserAuthData) error {
	if err := a.store.db.QueryRow(
		"INSERT INTO auth_data (user_id, mfcc) VALUES ($1, $2) RETURNING auth_data_id",
		authData.UserId,
		pq.Array(authData.Features),
	).Scan(&authData.AuthDataId); err != nil {
		return err
	}

	return nil
}

func (a *AuthDataRepository) All() ([]*auth.UserAuthData, error) {
	return nil, nil
}

func (a *AuthDataRepository) SaveMFCC(files []multipart.File, url string, userId int) {
	for _, file := range files {
		go a.CreateAuthData(file, url, userId)
	}
}

func (a *AuthDataRepository) CreateAuthData(file multipart.File, url string, userId int) {
	mfcc, err := auth.GetMFCCFeatures(file, url)
	if err != nil {
		panic(err)
	}

	ud := &auth.UserAuthData {
		UserId: userId,
		Features: mfcc,
	}

	if err := a.Create(ud); err != nil {
		panic(err)
	}
}
