package apiserver

import (
	"encoding/json"
	"github.com/reqww/go-rest-api/internal/app/grpcserver"
	"io"

	// 	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ctxKey int8

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.SetRequestID)
	s.router.Use(s.SetUpCors)
// 	s.router.Use(
// 		handlers.CORS(
// 			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
// 			handlers.AllowedHeaders([]string{"Token, Content-type"}),
// 			handlers.AllowedOrigins([]string{"localhost"}),
// 			handlers.AllowCredentials(),
// 		),
// 	)

	s.router.HandleFunc("/register", s.HandleUsersCreate()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/jwt", s.HandleJWTCreate()).Methods("POST", "OPTIONS")

	private := s.router.PathPrefix("/api").Subrouter()
	private.Use(s.AuthenticateUser)
	private.HandleFunc("/me", s.HandleMe()).Methods("GET", "OPTIONS")
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := s.ParseFiles(w, r, FilesCount)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var filesBytes [][]byte
		for _, file := range files {
			fileBytes, err := io.ReadAll(file)
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
			}
			filesBytes = append(filesBytes, fileBytes)
		}

		status, err := grpcserver.CreateUser(r.FormValue("email"), filesBytes)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusCreated, status)
	}
}

func (s *server) HandleJWTCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		files, err := s.ParseFiles(w, r, 1)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		file := files[0]
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		jwt, err := grpcserver.CreateJWT(r.FormValue("email"), fileBytes)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		var data = make(map[string]string)
		data["access"] = jwt

		s.respond(w, r, http.StatusOK, data)
	}
}

func (s *server) HandleMe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	s.logger.Infof("New response. METHOD = %v; URI = %v; CODE = %v", r.Method, r.URL, code)
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
