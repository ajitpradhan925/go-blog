// blog_handler.go

package handlers

import (
	"blog/internal/app/database"
	"blog/internal/app/model"
	"blog/internal/app/routes/middleware"
	"blog/internal/app/routes/utils"
	"blog/internal/app/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBlogs(c *gin.Context) {
	var blogs []model.Blog
	database.DB.Find(&blogs)
	c.JSON(http.StatusOK, blogs)
}

func CreateBlog(c *gin.Context) {
	var input validations.CreateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "Invalid input")
		return
	}
	userID := middleware.ExtractUserID(c)
	blog := model.Blog{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}

	database.DB.Create(&blog)
	c.JSON(http.StatusCreated, blog)
}
