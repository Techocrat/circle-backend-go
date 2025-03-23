package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
    var req SignUpRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := SignUpUser(req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, AuthResponse{Token: token})
}
