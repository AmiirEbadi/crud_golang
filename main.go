package main

import (
	"CRUD/controllers"
	"CRUD/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase() // new!

	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.GET("/posts", controllers.FindPosts)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.FindUsers)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run("localhost:8080")
}
