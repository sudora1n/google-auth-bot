package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("can't initialize zap logger")
	}
	Logger = logger.Sugar()
}
