package routes

import (
	"github.com/SampritiMitra/golang_apis/controllers"
	"github.com/gorilla/mux"
)

func Route_call(router *mux.Router){
	router.HandleFunc("/health", controllers.HomeLink)
	router.HandleFunc("/downloads/{id}", controllers.Status).Methods("GET")
	router.HandleFunc("/downloads", controllers.Download).Methods("POST")
}