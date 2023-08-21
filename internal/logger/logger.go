package logger

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	Logger = logger.Sugar()
}

func CloseLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
