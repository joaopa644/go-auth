package server

import (
	"go-auth-module/src/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	r := gin.Default()

	routes.ConfigureRoutes(r)

	r.Run(":3000")
}
