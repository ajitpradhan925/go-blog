// main.go
package main

import (
	"blog/internal/app/config"
	"blog/internal/app/database"
	"blog/internal/app/model"
	"blog/internal/app/server"
	middlewares "blog/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDatabase()

	err := database.DB.AutoMigrate(&model.User{}, &model.Blog{})
	if err != nil {
		panic("Failed to migrate tables")
	}

	r := server.SetupServer()

	// Use the global error handler middleware
	r.Use(middlewares.GlobalErrorHandler)

	// Add a custom 404 Not Found handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})

	r.Run(":8080")
}
