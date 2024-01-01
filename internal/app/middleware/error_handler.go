package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GlobalErrorHandler is a global middleware to handle errors
func GlobalErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// Log the error or perform any other actions you need
			fmt.Println("Recovered from panic:", err)

			// Return a JSON response with a 500 status code
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}()

	c.Next()
}
