package main

import (
	"github.com/Techocrat/circle-backend-go/config"
	"github.com/Techocrat/circle-backend-go/database"
	"github.com/Techocrat/circle-backend-go/internal/models"
)

func main() {
	config.LoadEnv()
	database.Connect()

	database.DB.AutoMigrate(&models.User{}, &models.Post{})
	// Set up routes here
}
