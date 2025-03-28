package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/Techocrat/circle-backend-go/database"
	"github.com/Techocrat/circle-backend-go/internal/models"
)

func SignUpUser(req SignUpRequest) (string, error) {
	// Check if user already exists
	var existing models.User
	result := database.DB.Where("email = ?", req.Email).First(&existing)
	if result.Error == nil {
		return "", errors.New("user with this email already exists")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return "", err
	}

	// Generate token
	return GenerateJWT(user.ID)
}

func SignInUser(req SignInRequest) (string, error) {
	var user models.User
	err := database.DB.Where("email = ? OR username = ?", req.Identifier, req.Identifier).First(&user).Error
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
