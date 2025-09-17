package utils

import (
	"fmt"
	"os"
	"time"

	"main/models"
	"main/structs"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken generates a JWT token for a given user.
// The token is signed with the HS256 algorithm and contains the user's information.
// The token is valid for 2 hours from the moment it is generated.
// If an error occurs while generating the token, it is returned as the second argument.
//
// @param user models.PublicUser
// @return string, error
func GenerateToken(user models.PublicUser) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)

	claims := &structs.UserClaims{
		PublicUser: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gdp-nexus",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error al firmar el token: %w", err)
	}

	return tokenString, nil
}
