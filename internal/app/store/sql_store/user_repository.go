package sql_store

import (
	"database/sql"
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) All() ([]*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindById(userID int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email FROM users WHERE user_id = $1", userID,
	).Scan(&u.UserId, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email) VALUES ($1) RETURNING user_id",
		u.Email,
	).Scan(&u.UserId); err != nil {
		return err
	}

	return nil
}


func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email FROM users WHERE email = $1", email,
		).Scan(&u.UserId, &u.Email); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
		return nil, err
	}

	return u, nil
}