package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	connectDB()
	router := gin.Default()
	users := router.Group("/users")
	{
		users.GET("", getUsers)
		users.POST("", postUser)
		users.GET("/:id", getUserByID)

	}
	posts := router.Group("/posts")
	{
		posts.GET("", getPosts)
		posts.POST("", createPost)
		posts.GET("/:id", getPostsByUserId)

	}

	router.Run("localhost:8080")
}
