package sql_store

import (
	"github.com/reqww/go-rest-api/internal/app/auth"
	"mime/multipart"
)

type AuthDataRepository struct {
	store *Store
}

func (a *AuthDataRepository) Create(authData *auth.UserAuthData) error {

	if err := a.store.db.QueryRow(
		"INSERT INTO auth_data (user_id, features) VALUES ($1, $2) RETURNING auth_data_id",
		authData.UserId,
		authData.Features,
	).Scan(&authData.AuthDataId); err != nil {
		return err
	}

	return nil
}

func (a *AuthDataRepository) All() ([]*auth.UserAuthData, error) {
	rows, err := a.store.db.Query("SELECT features, user_id FROM auth_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*auth.UserAuthData

	for rows.Next() {
		authData := &auth.UserAuthData{}
		if err := rows.Scan(&authData.Features, &authData.UserId); err != nil {
			return nil, err
		}
		res = append(res, authData)
	}

	return res, nil
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

	features := new(auth.DataFeatures)
	features.MFCC = mfcc

	ud := &auth.UserAuthData {
		UserId: userId,
		Features: *features,
	}

	if err := a.Create(ud); err != nil {
		panic(err)
	}
}

func (a *AuthDataRepository) DetermineUserBySound(mfcc [][]float64) (int, error) {

	authData, err := a.All()
	if err != nil {
		return 0, err
	}

	m := auth.NewMap()

	for _, data := range authData {
		difference := data.Features.CalculateDifference(mfcc)
		m.AddToKey(data.UserId, difference)
	}

	return m.UserMin(), nil
}
