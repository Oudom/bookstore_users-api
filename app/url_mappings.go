package app

import (
	"github.com/Oudom/bookstore_users-api/controllers/ping"
	"github.com/Oudom/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUer)
	router.PATCH("/users/:user_id", users.UpdateUer)
	// router.GET("/users/search", users.SearchUser)

}
