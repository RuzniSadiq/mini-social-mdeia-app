package auth

import (
	"net/http"

	"mini-social-media/api/generics"

	"github.com/gin-gonic/gin"
)

// LoginController is responsible for handling requests related to authentication
func Login(c *gin.Context) {
	var credentials UserCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		generics.CustomJsonResponseError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	token, err := LoginService(credentials)
	if err != nil {
		generics.CustomJsonResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK,"token": token})
}
