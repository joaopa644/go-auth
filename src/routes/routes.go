package routes

import (
	helloworld "go-auth-module/src/controllers/hello-world"
)

func ConfigureRoutes() {
	helloworld.HelloWord()
}
