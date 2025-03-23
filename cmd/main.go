package main

import (
	"log"
	"os"

	"github.com/Techocrat/circle-backend-go/database"
	"github.com/Techocrat/circle-backend-go/internal/models"
	"github.com/Techocrat/circle-backend-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		routes.AuthRoutes(api)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run(":8080")
}
