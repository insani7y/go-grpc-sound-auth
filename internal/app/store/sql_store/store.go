package sql_store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/reqww/go-rest-api/internal/app/store"
)

type Store struct {
	db *sql.DB
	userRepository *UserRepository
	authDataRepository *AuthDataRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) AuthData() store.AuthDataRepository {
	if s.authDataRepository != nil {
		return s.authDataRepository
	}

	s.authDataRepository = &AuthDataRepository{
		store: s,
	}

	return s.authDataRepository
}
