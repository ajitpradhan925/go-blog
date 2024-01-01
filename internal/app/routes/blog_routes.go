// blog_routes.go
package routes

import (
	"blog/internal/app/handlers"
	"blog/internal/app/routes/middleware"

	"github.com/gin-gonic/gin"
)

func SetupBlogRoutes(r *gin.RouterGroup) {
	blogRoutes := r.Group("/blogs")
	{
		blogRoutes.GET("/", handlers.GetAllBlogs)
		blogRoutes.POST("/", middleware.AuthMiddleware(), handlers.CreateBlog)
	}
}
