package auth

type SignUpRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct {
    Token string `json:"token"`
}