// database.go - containers the database connection
package database

import (
	"blog/internal/app/config"
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := config.DatabaseURL

	fmt.Println("Database DSN:", config.DatabaseURL) // Print the DSN

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	DB = db
}
