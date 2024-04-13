package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hamillka/avitoTech24/internal/handlers/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func Router(bs BannerService, logger *zap.SugaredLogger) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	bh := NewBannerHandler(bs, logger)
	router.HandleFunc("/banner", bh.GetBanners).Methods("GET")
	router.HandleFunc("/banner", bh.CreateBanner).Methods("POST")

	router.HandleFunc("/banner/{id}", bh.UpdateBanner).Methods("PATCH")
	router.HandleFunc("/banner/{id}", bh.DeleteBanner).Methods("DELETE")

	router.HandleFunc("/user_banner", bh.GetUserBanner).Methods("GET")

	router.Use(middlewares.AuthMiddleware)
	return router
}
