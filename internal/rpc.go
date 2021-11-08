package internal

import (
	tokenV1 "account-server/pb/basic/token/v1"
	userV1 "account-server/pb/basic/user/v1"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var userServiceClient userV1.UserServiceClient
var tokenServiceClient tokenV1.TokenServiceClient

func NewUserServiceClient() userV1.UserServiceClient {
	if nil == userServiceClient {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", Config.UserServer.Host, Config.UserServer.Port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		userServiceClient = userV1.NewUserServiceClient(conn)
	}
	return userServiceClient
}

func NewTokenServiceClient() tokenV1.TokenServiceClient {
	if nil == tokenServiceClient {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", Config.TokenServer.Host, Config.TokenServer.Port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		tokenServiceClient = tokenV1.NewTokenServiceClient(conn)
	}
	return tokenServiceClient
}
