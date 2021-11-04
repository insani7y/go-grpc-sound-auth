package test_store

import (
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	u.UserId = len(r.users) + 1
	r.users[u.UserId] = u
	return nil
}

func (r *UserRepository) FindById(userID int) (*model.User, error) {
	u, ok := r.users[userID]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) All() ([]*model.User, error) {
	return nil, nil
}
