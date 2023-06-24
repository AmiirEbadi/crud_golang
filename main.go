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
	router.Run("localhost:8080")
}
