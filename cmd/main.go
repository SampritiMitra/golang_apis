package main

import (
	"github.com/SampritiMitra/golang_apis/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.Route_call(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
