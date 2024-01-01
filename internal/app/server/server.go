// internal/app/server/server.go
package server

import (
	"blog/internal/app/routes"

	"github.com/gin-gonic/gin"
)

// SetupServer configures and returns a Gin server
func SetupServer() *gin.Engine {
	r := gin.Default()

	// Setup routes
	apiRoutes := r.Group("/api")
	{
		routes.SetupBlogRoutes(apiRoutes)
		routes.SetupUserRoutes(apiRoutes)
	}

	return r
}
