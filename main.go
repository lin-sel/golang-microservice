package main

import (
	"log"
	"net/http"

	"bitbucket.com/lin-sel/golang-microservice/controller"
	"bitbucket.com/lin-sel/golang-microservice/service"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	ser := service.NewContactService()
	con := controller.NewContactController(ser)
	con.RouteSpecifier(route)

	// Server Launched
	server := http.Server{
		Handler: route,
		Addr:    "0.0.0.0:8080",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server has Stop due to %s", err.Error())
	}
}
