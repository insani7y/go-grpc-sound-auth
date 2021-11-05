package auth

import (
	"bytes"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"io"
	"log"
	"math"
	"mime/multipart"
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

type MFCCResponse struct {
	MFCC [][]float64 `json:"mfcc"`
}

func GetMFCCFeatures(file multipart.File, url string) ([][]float64, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	mfccRes := &MFCCResponse{}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", "sound.wav")

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, file)

	if err != nil {
		return nil, err
	}

	writer.Close()
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}

	if err := json.NewDecoder(rsp.Body).Decode(mfccRes); err != nil {
		return nil, err
	}

	return mfccRes.MFCC, nil
}

func CalculateVecAbs(vec []float64) float64 {
	var res float64
	for _, value := range vec {
		res += math.Pow(value, float64(2))
	}

	return math.Sqrt(res)
}