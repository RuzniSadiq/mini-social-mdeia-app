package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerHost string
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var AppEnvConfig EnvConfig

// LoadEnvConfig loads environment variables from the .env file
func LoadEnvConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	AppEnvConfig = EnvConfig{
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "social_media"),
		JWTSecret:  getEnv("JWT_SECRET", "my_jwt_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
