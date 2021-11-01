package apiserver

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/reqww/go-rest-api/internal/app/auth"
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ctxKey int8

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.SetRequestID)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/register", s.HandleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/jwt", s.HandleJWTCreate()).Methods("POST")

	private := s.router.PathPrefix("/api").Subrouter()
	private.Use(s.AuthenticateUser)
	private.HandleFunc("/me", s.HandleMe()).Methods("GET")
}

func (s *server) HandleUsersCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		filesBytes, err := s.ParseFiles(w, r)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User {
			Email: r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := u.GetAllThingsDone(filesBytes); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) HandleJWTCreate() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		token, err := auth.GenerateJWT(u.ID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		data := make(map[string]string)
		data["access"] = token

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