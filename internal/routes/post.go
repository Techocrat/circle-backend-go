package routes

import (
	"github.com/Techocrat/circle-backend-go/internal/handlers"
	"github.com/Techocrat/circle-backend-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(rg *gin.RouterGroup) {
	post := rg.Group("/posts")
	post.Use(middleware.JWTAuthMiddleware())
	post.POST("", handlers.CreatePost)
}
