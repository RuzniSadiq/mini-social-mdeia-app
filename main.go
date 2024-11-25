package main

import (
	"fmt"
	"log"
	"mini-social-media/api"
	"mini-social-media/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvConfig()
	initializers.InitDB()
}

func main() {
	// Initialize router with API versioning
	r := gin.Default()
	api.SetupRoutes(r)

	server := fmt.Sprintf("%s:%s", initializers.AppEnvConfig.ServerHost, initializers.AppEnvConfig.ServerPort)

	// Start the server
	if err := r.Run(server); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
