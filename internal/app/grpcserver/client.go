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
		return "", err
	}

	c := NewSoundAuthServiceClient(conn)

	res, err := c.CreateJWT(context.Background(), &CreateJWTMessage{
		Email:         email,
		File1:         fileBytes,
	})
	if err != nil {
		return "", err
	}

	return res.GetAccess(), nil
}

func CreateUser(email string, files [][]byte) (int64, error) {
	config := NewConfig()

	conn, err := grpc.Dial(config.BindGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := NewSoundAuthServiceClient(conn)

	res, err := c.CreateUser(context.Background(), &UserCreateMessage{
		Email:         email,
		File1:         files[0],
		File2:         files[1],
		File3:         files[2],
		File4:         files[3],
		File5:         files[4],
	})

	if err != nil {
		return 1, err
	}

	return res.GetStatus(), nil
}
