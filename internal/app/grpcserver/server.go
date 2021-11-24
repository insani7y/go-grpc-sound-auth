package grpcserver

import (
	"github.com/reqww/go-rest-api/internal/app/store"
)

func newGRPCServer(store store.Store) *GRPCServer {
	s := &GRPCServer{
		store:  store,
	}

	return s
}
