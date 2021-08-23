package server

import (
	"go-auth-module/src/routes"
	"log"
	"net/http"
)

func StartServer() {

	routes.ConfigureRoutes()

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
