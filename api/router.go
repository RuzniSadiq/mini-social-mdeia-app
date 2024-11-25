package api

import (
	"mini-social-media/api/auth"
	"mini-social-media/api/post_management"
	"mini-social-media/api/user_management"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth.RegisterRoutes(v1)
		post_management.RegisterRoutes(v1)
		user_management.RegisterRoutes(v1)
	}
}
