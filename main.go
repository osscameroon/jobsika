package main

import (
	"github.com/elhmn/camerdevs/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Health check endpoint
	router.GET("/health", handlers.Health)

	//Authentication endpoints
	router.POST("/signin", handlers.Signin)
	router.POST("/signup", handlers.Signup)

	//Endpoints that require authorisation
	authorized := router.Group("/")
	authorized.Use(handlers.AuthMiddleware())
	{
		// 		router.GET("/collections/", handlers.GetCollections)
		// 		router.POST("/collections/", handlers.SetCollections)
	}

	if err := router.Run(":7000"); err != nil {
		return
	}
}
