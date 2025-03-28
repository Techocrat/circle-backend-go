package routes

import (
	"github.com/Techocrat/circle-backend-go/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignUp)
		authGroup.POST("/signin", auth.SignIn)
	}
}
