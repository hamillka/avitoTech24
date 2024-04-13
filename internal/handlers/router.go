package handlers

import (
	"github.com/gorilla/mux"
	_ "github.com/hamillka/avitoTech24/api"
	"github.com/hamillka/avitoTech24/internal/handlers/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func Router(bs BannerService, logger *zap.SugaredLogger) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	api := router.PathPrefix("").Subrouter()

	bh := NewBannerHandler(bs, logger)
	api.HandleFunc("/banner", bh.GetBanners).Subrouter().Methods("GET")
	api.HandleFunc("/banner", bh.CreateBanner).Subrouter().Methods("POST")

	api.HandleFunc("/banner/{id}", bh.UpdateBanner).Subrouter().Methods("PATCH")
	api.HandleFunc("/banner/{id}", bh.DeleteBanner).Subrouter().Methods("DELETE")

	api.HandleFunc("/user_banner", bh.GetUserBanner).Subrouter().Methods("GET")
	api.Use(middlewares.AuthMiddleware)
	return router
}
