package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Techocrat/circle-backend-go/database"
	"github.com/Techocrat/circle-backend-go/internal/models"
	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url"`
}

func CreatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	log.Printf("UserID: %v", userID)
	
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		UserID:      userID.(uint),
		Description: req.Description,
		ImageURL:    req.ImageURL,
		CreatedAt:   time.Now(),
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}
