package main

import (
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/config"
	database "github.com/sudora1n/google-auth-bot/internal/microservice-api/database"
	grpcserver "github.com/sudora1n/google-auth-bot/internal/microservice-api/grpc-server"
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/logger"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	orm, err := database.InitOrm()
	if err != nil {
		logger.Logger.Fatalf("can't initialize orm %v", err)
	}

	logger.Logger.Info("Starting microservice-api")

	err = grpcserver.Start(orm)
	if err != nil {
		logger.Logger.Fatalf("can't start grpc server %v", err)
	}
}
