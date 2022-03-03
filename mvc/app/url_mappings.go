package app

import (
	"github.com/manyiu/go-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}