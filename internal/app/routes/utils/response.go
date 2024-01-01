// response.go
package utils

import "github.com/gin-gonic/gin"

func BadRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{"error": message})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(404, gin.H{"error": message})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(401, gin.H{"error": message})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(500, gin.H{"error": message})
}
