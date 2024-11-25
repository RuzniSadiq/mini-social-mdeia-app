package post_management

import (
	"mini-social-media/middlewares"

	"github.com/gin-gonic/gin"
)


var authMiddleware = middlewares.AuthMiddleware()

func RegisterRoutes(r *gin.RouterGroup) {
	postController := NewPostController()
	postGroup := r.Group("/posts") 
	{
		postGroup.POST("/", authMiddleware,postController.Create)
		postGroup.GET("/:id", authMiddleware,postController.Retrieve)
		postGroup.GET("/", authMiddleware,postController.List)
		postGroup.PUT("/:id", authMiddleware,postController.Update)
		postGroup.DELETE("/:id", authMiddleware,postController.Delete)

	}

	commentController := NewCommentController()
	commentGroup := r.Group("/comments") 
	{
		commentGroup.POST("/", authMiddleware,commentController.Create)
		commentGroup.GET("/:id", authMiddleware,commentController.Retrieve)
		commentGroup.GET("/", authMiddleware,commentController.List)
		commentGroup.PUT("/:id", authMiddleware,commentController.Update)
		commentGroup.DELETE("/:id", authMiddleware,commentController.Delete)

	}

	likeController := NewLikeController()
	likeGroup := r.Group("/likes") 
	{
		likeGroup.POST("/", authMiddleware,likeController.Create)
		likeGroup.GET("/:id", authMiddleware,likeController.Retrieve)
		likeGroup.GET("/", authMiddleware,likeController.List)
		likeGroup.DELETE("/:id", authMiddleware,likeController.Delete)

	}
}
