// internal/app/routes/utils/token.go
package utils

import (
	"blog/internal/app/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.SecretKey)

// Claims struct to store information in the JWT token
type Claims struct {
	UserID uint `json:"id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the given user ID
func GenerateToken(userID uint) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
