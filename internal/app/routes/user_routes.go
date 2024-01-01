// user_routes.go
package routes

import (
	"blog/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", handlers.Register)
		userRoutes.POST("/login", handlers.Login)
	}
}
