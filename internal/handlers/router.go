package handlers

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
