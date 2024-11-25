package middlewares

import (
	"net/http"
	"strings"

	"mini-social-media/api/auth"

	"mini-social-media/api/generics"

	"github.com/gin-gonic/gin"

	"mini-social-media/initializers"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			generics.CustomJsonResponseError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := auth.ParseToken(token, initializers.AppEnvConfig.JWTSecret)
		if err != nil {
			generics.CustomJsonResponseError(c, http.StatusUnauthorized, "Invalid Token")
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
