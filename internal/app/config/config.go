// config.go - contains all config variables
package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	DatabaseURL = ""
	SecretKey   = ""
)

// LoadConfig loads configuration from environment variables or config files
func LoadConfig() {
	viper.SetConfigFile(".env") // You can use a different config file format if needed, like YAML or JSON
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	DatabaseURL = viper.GetString("DATABASE_URL")
	SecretKey = viper.GetString("SECRET_KEY")
}
