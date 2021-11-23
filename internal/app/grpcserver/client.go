package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"log"
)


func CreateJWT(email string, fileBytes []byte) (string, error) {
	config := NewConfig()

	conn, err := grpc.Dial(config.BindGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := NewSoundAuthServiceClient(conn)

	res, err := c.CreateJWT(context.Background(), &CreateJWTMessage{
		Email:         email,
		File1:         fileBytes,
	})
	if err != nil {
		log.Fatal(err)
	}

	return res.GetAccess(), nil
}

func CreateUser() (int64, error) {
	config := NewConfig()

	conn, err := grpc.Dial(config.BindGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := NewSoundAuthServiceClient(conn)

	res, err := c.CreateUser(context.Background(), &UserCreateMessage{
		Email:         "",
		File1:         nil,
		File2:         nil,
		File3:         nil,
		File4:         nil,
		File5:         nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	return res.GetStatus(), nil
}
