package user_management

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", RegisterUserController)
		userGroup.DELETE("/delete", UserDeleteController)
	}
}
