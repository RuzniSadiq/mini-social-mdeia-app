package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
    dbHost := AppEnvConfig.DBHost
    dbPort := AppEnvConfig.DBPort
    dbUser := AppEnvConfig.DBUser
    dbPassword := AppEnvConfig.DBPassword
    dbName := AppEnvConfig.DBName

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    fmt.Println("Database connected successfully")
}
