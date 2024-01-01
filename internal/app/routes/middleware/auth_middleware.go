// auth_middleware.go - contains middleware to extract user info from JWT Token
package middleware

import (
	"blog/internal/app/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt") // assuming you are storing the JWT in a cookie
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the key or secret used to sign the token here
			return []byte(config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

func ExtractUserID(c *gin.Context) uint {
	tokenString, err := c.Cookie("jwt") // assuming you are storing the JWT in a cookie
	if err != nil {
		return 0
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the key or secret used to sign the token here
		return []byte(config.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}

	// Assuming your user ID is stored as "user_id" in the JWT claims
	userID := uint(claims["id"].(float64))
	return userID
}
