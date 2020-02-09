package app

import (
	"BicepsFavor/bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.FindUser)
	router.POST("/users", users.CreateUser)
}
