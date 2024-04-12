package main

import (
	"net/http"

	"github.com/hamillka/avitoTech24/internal/config"
	"github.com/hamillka/avitoTech24/internal/db"
	"github.com/hamillka/avitoTech24/internal/handlers"
	"github.com/hamillka/avitoTech24/internal/logger"
	"github.com/hamillka/avitoTech24/internal/repositories"
	"github.com/hamillka/avitoTech24/internal/services"
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

	db, err := db.CreateConnection(&config.DB)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Errorf("Error while closing connection to db: %v", err)
		}
	}()

	if err != nil {
		logger.Fatalf("Error while connecting to database: %v", err)
	}

	br := repositories.NewBannerRepository(db)
	fr := repositories.NewFeatureRepository(db)
	tr := repositories.NewTagRepository(db)
	btr := repositories.NewBannerTagRepository(db)

	bs := services.NewBannerService(br, fr, tr, btr)

	r := handlers.Router(bs, logger)

	port := config.Port
	logger.Info("Server is started on port ", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.Fatalf("Error while starting server: %v", err)
	}
}
