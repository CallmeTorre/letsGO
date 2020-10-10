package app

import (
	"github.com/CallmeTorre/letsGO/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
