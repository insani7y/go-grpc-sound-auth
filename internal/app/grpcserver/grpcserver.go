package grpcserver

import (
	"context"
	"fmt"
	"github.com/reqww/go-rest-api/internal/app/auth"
	"github.com/reqww/go-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	store  store.Store
}

func Start(config *Config) error {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	RegisterSoundAuthServiceServer(s, srv)

	l, err := net.Listen("tcp", config.BindGRPCAddr)

	if err != nil {
		log.Fatal(err)
	}

	logrus.Info(fmt.Sprintf("GRPC server started at %v", config.BindGRPCAddr))

	return s.Serve(l)
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *UserCreateMessage) (*Status, error) {
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
