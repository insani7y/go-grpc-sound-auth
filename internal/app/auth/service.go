package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/reqww/go-rest-api/internal/app/fileHandler"
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

func GetMFCCFeatures(filesBytes [][]byte) ([][]float64, error) {

	handler := fileHandler.New()

	var res [][]float64

	for _, fileBytes := range filesBytes {
		data, err := handler.DetermineAmplitudeValues(fileBytes)
		if err != nil {
			return nil, err
		}

		frames := handler.FrameCut(data)

		amplitudes := handler.FourierTransform(frames)

		melAmplitudes := handler.ToMelScale(amplitudes)

		features := handler.GetMelFeaturesArr(melAmplitudes)

		res = append(res, features)
	}

	return res, nil
}
