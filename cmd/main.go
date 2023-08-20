package main

import (
	"github.com/franciscofferraz/go-struct/internal/app"
	"github.com/franciscofferraz/go-struct/internal/config"
	"github.com/franciscofferraz/go-struct/internal/database"
	"github.com/franciscofferraz/go-struct/internal/server"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	app.InitLogger()
	defer app.CloseLogger()

	err := godotenv.Load()
	if err != nil {
		app.Logger.Fatal("Error loading .env file", zap.Error(err))
	}

	cfg, err := config.ParseFlags()
	if err != nil {
		app.Logger.Fatal("Failed to parse command-line flags", zap.Error(err))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		app.Logger.Fatal("Failed to connect to the database", zap.Error(err))
		panic(err)
	}
	defer db.Close()

	server := server.NewServer(app.Logger, db, cfg)

	err = server.Start(":8080")
	if err != nil {
		app.Logger.Fatal("Failed to start the server", zap.Error(err))
	}
}
