package auth

import "github.com/gin-gonic/gin"

// RegisterRoutes sets up the routes for the authentication module
func RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", Login)
	}
}
