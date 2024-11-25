package user_management

import (
	"errors"
	"mini-social-media/initializers"

	"mini-social-media/api/generics"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct {
	*generics.GenericService[User]
}

func NewUserService() *UserService {
	return &UserService{
		GenericService: generics.NewGenericService[User](initializers.DB),
	}
}

func (s *UserService) Delete(c *gin.Context, id any) error {
	return s.DeleteService(c, id)

}

func (s *UserService) CheckExistingUser(username string) (*User, error) {
	var existingUser User
	if err := initializers.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}
	return nil, nil
}

// RegisterUser creates a new user
func (s *UserService) RegisterUserService(username, password string) (*User, error) {
	

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the new user
	user := User{Username: username, Password: string(hashedPassword)}
	if err := initializers.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
