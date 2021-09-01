package helloworld

import (
	"net/http"

	"github.com/gin-gonic/gin"

	helloworldrequestmodel "go-auth-module/src/model/request/hello-world"
)

func HelloWord(r *gin.Engine) {
	r.GET("/hello-world", func(c *gin.Context) {

		var message helloworldrequestmodel.HelloWorldRequestModel

		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"response": message,
		})
	})
}

func LoadRoutes(r *gin.Engine) {
	HelloWord(r)
}
