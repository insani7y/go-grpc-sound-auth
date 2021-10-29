package auth

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func GenerateJWT(userID int) (string, error) {
	config := NewConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		userID,
	})

	tokenString, err := token.SignedString([]byte(config.AuthSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsAuthenticated(r *http.Request) (*jwt.Token, error) {
	config := NewConfig()

	if access := r.Header[config.TokenHeader]; access != nil {
		token, err := jwt.Parse(access[0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrUnauthorized
			}
			return []byte(config.AuthSecret), nil
		})
		if err != nil {
			return nil, err
		}
		if token.Valid {
			return token, nil
		}

		return nil, InvalidToken

	} else {
		return nil, ErrUnauthorized
	}
}
