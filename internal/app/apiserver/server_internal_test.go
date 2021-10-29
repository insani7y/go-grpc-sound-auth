package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/reqww/go-rest-api/internal/app/auth"
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store/test_store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(test_store.New())

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/register", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := test_store.New()
	store.User().Create(u)
	s := newServer(store)

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"email": "user1@example.org",
				"password": "password",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/jwt", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_AuthenticateUser(t *testing.T) {
	store := test_store.New()
	u := model.TestUser(t)
	store.User().Create(u)

	token, _ := auth.GenerateJWT(u.ID)

	testCases := []struct{
		name string
		token string
		expectedCode int
	}{
		{
			name: "authenticated",
			token: token,
			expectedCode: http.StatusOK,
		},
		{
			name: "unAuthenticated",
			token: "",
			expectedCode: http.StatusUnauthorized,
		},
	}

	s := newServer(store)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	config := auth.NewConfig()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(config.TokenHeader, tc.token)
			s.AuthenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}