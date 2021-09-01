package routes

import (
	helloworld "go-auth-module/src/controllers/hello-world"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine) {
	helloworld.LoadRoutes(r)
}
