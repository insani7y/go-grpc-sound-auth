package auth

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	InvalidToken = errors.New("invalid token")
)
