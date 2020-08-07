package config

import (
	"os"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Print("Error loading .env file")
	// }
	return os.Getenv(key)
}
