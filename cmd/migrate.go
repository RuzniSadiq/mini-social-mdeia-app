package main

import (
	"fmt"
	"log"
	"mini-social-media/api/post_management"
	"mini-social-media/api/user_management"
	"mini-social-media/initializers"
)

func init() {
	// Load the environment variables
	initializers.LoadEnvConfig()
	// Initialize the database connection
	initializers.InitDB()
}

func RunMigrations() {
	// Initialize the DB connection
	db := initializers.DB
	
	// Perform automigrations
	err := db.AutoMigrate(
		&user_management.User{},
		&post_management.Post{},
		&post_management.Comment{},
		&post_management.Like{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	fmt.Println("Database migrated successfully")
}

func main() {
	RunMigrations()
}
