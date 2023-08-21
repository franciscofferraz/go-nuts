package main

import (
	"github.com/franciscofferraz/go-struct/internal/api"
	"github.com/franciscofferraz/go-struct/internal/config"
	"github.com/franciscofferraz/go-struct/internal/db"
	"github.com/franciscofferraz/go-struct/internal/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file", zap.Error(err))
	}

	cfg := config.NewConfig()

	err = cfg.ParseFlags()
	if err != nil {
		logger.Logger.Fatal("Failed to parse command-line flags", zap.Error(err))
	}

	db, err := db.Connect(cfg)
	if err != nil {
		logger.Logger.Fatal("Failed to connect to the database", zap.Error(err))
		panic(err)
	}
	defer db.Close()

	h, _ := cfg.InitializeHandlers(db)
	srv := api.NewAPI(logger.Logger, cfg, h)

	err = srv.Run()
	if err != nil {
		logger.Logger.Fatal("Failed to start the server", zap.Error(err))
	}
}
