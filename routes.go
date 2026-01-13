package main

import (
	"gin-blueprint/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", handlers.GetAllUsers)
			users.GET("/:id", handlers.GetUser)
			users.POST("", handlers.CreateUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
			users.GET("/:id/posts", handlers.GetUserPosts)
		}

		posts := v1.Group("/posts")
		{
			posts.GET("", handlers.GetAllPosts)
			posts.GET("/:id", handlers.GetPost)
			posts.POST("", handlers.CreatePost)
		}
	}
}
