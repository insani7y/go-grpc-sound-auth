package grpcserver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/reqww/go-rest-api/internal/app/auth"
	"github.com/reqww/go-rest-api/internal/app/model"
	"github.com/reqww/go-rest-api/internal/app/store"
	"github.com/reqww/go-rest-api/internal/app/store/sql_store"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	store  store.Store
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Start(config *Config) error {
	s := grpc.NewServer()

	db, err := newDB(config.DatabaseURL)
	defer db.Close()

	if err != nil {
		return err
	}
	store := sql_store.New(db)

	srv := newGRPCServer(store)

	RegisterSoundAuthServiceServer(s, srv)

	l, err := net.Listen("tcp", config.BindGRPCAddr)

	if err != nil {
		log.Fatal(err)
	}

	logrus.Info(fmt.Sprintf("GRPC server started at %v", config.BindGRPCAddr))

	return s.Serve(l)
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *UserCreateMessage) (*Status, error) {
	config := auth.NewConfig()

	u := &model.User{
		Email: req.Email,
	}

	if err := s.store.User().Create(u); err != nil {
		return nil, err
	}

	files := [][]byte{req.File1, req.File2, req.File3, req.File4, req.File5}

	s.store.AuthData().SaveMFCC(files, config.MFCCUrl, u.UserId)

	return &Status{Status: 0}, nil
}

func (s *GRPCServer) CreateJWT(ctx context.Context, req *CreateJWTMessage) (*JWT, error) {
	config := auth.NewConfig()

	mfcc, err := auth.GetMFCCFeatures(req.GetFile1(), config.MFCCUrl)

	userId, err := s.store.AuthData().DetermineUserBySound(mfcc)

	if err != nil {
		return nil, err
	}

	email := req.Email
	u, err := s.store.User().FindByEmail(email)

	if err != nil || u.UserId != userId {
		return nil, err
	}

	token, err := auth.GenerateJWT(userId)
	if err != nil {
		return nil, err
	}

	return &JWT{Access: token}, nil
}
