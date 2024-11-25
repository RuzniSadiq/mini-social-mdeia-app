package auth

import (
	"errors"

	"mini-social-media/initializers"

	"mini-social-media/api/user_management"

	"golang.org/x/crypto/bcrypt"
)
// LoginService is responsible for authenticating a user
func LoginService(credentials UserCredentials) (string, error) {
	var user user_management.User
	if err := initializers.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	
	token, err := GenerateToken(int(user.ID), initializers.AppEnvConfig.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
