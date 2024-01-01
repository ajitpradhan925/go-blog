// user_handler.go
package handlers

import (
	"blog/internal/app/database"
	"blog/internal/app/model"
	"blog/internal/app/routes/utils"
	"blog/internal/app/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input validations.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "Invalid input")
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var input validations.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "Invalid input")
		return
	}

	var user model.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.NotFound(c, "User not found")
		return
	}

	// Check password (you should use a secure password hashing library)
	if user.Password != input.Password {
		utils.Unauthorized(c, "Invalid password")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.InternalServerError(c, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
