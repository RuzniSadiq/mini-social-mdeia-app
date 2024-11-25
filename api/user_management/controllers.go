package user_management

import (
	"net/http"
	"strconv"

	"mini-social-media/api/generics"

	"github.com/gin-gonic/gin"
)

var userService = NewUserService()



func RegisterUserController(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		generics.CustomJsonResponseError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := userService.CheckExistingUser(request.Username)

	if err != nil {
		generics.CustomJsonResponseError(c, http.StatusBadRequest, "Username already exists")
		return
	}

	user, err := userService.RegisterUserService(request.Username, request.Password)
	if err != nil {
		generics.CustomJsonResponseError(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	generics.CustomJsonResponseData(c, http.StatusCreated, user)
	
}



func UserDeleteController(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		generics.CustomJsonResponseError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	// Perform deletion based on the model
	err = userService.DeleteService(c, id)

	// Handle errors from the service
	if err != nil {
		generics.CustomJsonResponseError(c, http.StatusForbidden, err.Error())
		return
	}

	generics.CustomJsonResponseData(c, http.StatusOK, "User deleted successfully")
}
