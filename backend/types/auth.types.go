package types

import (
	"main/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	NickName  string `json:"nickname" binding:"required"`
}

type UserClaims struct {
	models.PublicUser
	jwt.RegisteredClaims
}
