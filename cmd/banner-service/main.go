package main

import (
	"github.com/hamillka/avitoTech24/internal/config"
	"github.com/hamillka/avitoTech24/internal/db"
	"github.com/hamillka/avitoTech24/internal/logger"
)

func main() {
	config, err := config.New()
	logger := logger.CreateLogger(config.Log)

	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Errorf("Error while syncing logger: %v", err)
		}
	}()

	if err != nil {
		logger.Errorf("Something went wrong with config: %v", err)
	}

	db, err := db.CreateConnection(config.Db)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Errorf("Error while closing connection to db: %v", err)
		}
	}()

	if err != nil {
		logger.Fatalf("Error while connecting to database: %v", err)
	}
}

