package grpcserver

import (
	"net"

	"google.golang.org/grpc"

	orm "github.com/sudora1n/google-auth-bot/internal/microservice-api/database"
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/grpc-server/routes"
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/logger"
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/proto"
)

func Start(ORMFunctions *orm.ORMFunctions) error {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		logger.Logger.Fatalf("can't listen to port %v", err)
	}

	grpcServer := grpc.NewServer()

	ToTPServer := routes.ToTPServer{ORMFunctions: ORMFunctions}
	proto.RegisterToTPsServer(grpcServer, &ToTPServer)

	UserServer := routes.UserServer{ORMFunctions: ORMFunctions}
	proto.RegisterUsersServer(grpcServer, &UserServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Logger.Fatalf("can't start grpc server %v", err)
	}
	return nil
}
