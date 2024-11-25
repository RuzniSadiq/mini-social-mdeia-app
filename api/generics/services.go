package generics

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



// NewGenericService returns a new instance of GenericService for a given entity type T
func NewGenericService[T any](db *gorm.DB) *GenericService[T] {
	return &GenericService[T]{DB: db}
}

// CheckOwnership checks if the current user (from the context) is the owner of the entity
func (s *GenericService[T]) CheckOwnership(c *gin.Context, entity T) (bool, error) {
	// Retrieve user ID from the context
	userID, exists := c.Get("user_id")
	if !exists {
		return false, fmt.Errorf("user not authenticated")
	}

	// Query to fetch the UserID associated with this entity
	var userIDFromDb uint
	if err := s.DB.Model(&entity).Select("user_id").Scan(&userIDFromDb).Error; err != nil {
		return false, fmt.Errorf("could not retrieve user ID for the record")
	}

	// Check ownership by comparing the UserID from the database with the user ID from the context
	if int(userIDFromDb) != userID {
		return false, fmt.Errorf("unauthorized: user does not have permission to access this record")
	}

	return true, nil
}

// Create adds a new entity of type T to the database
func (s *GenericService[T]) CreateService(entity *T) (*T, error) {
	// Optionally, set the user ID in the entity (if required for certain entities)
	// entity.UserID = userID

	if err := s.DB.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// Retrieve retrieves an entity by ID, with optional preloads
func (s *GenericService[T]) RetrieveService(id any, preloads ...string) (*T, error) {
	var entity T
	query := s.DB
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// List retrieves all entities of type T, with optional preloads
func (s *GenericService[T]) ListService(preloads ...string) ([]T, error) {
	var results []T
    query := s.DB
    for _, preload := range preloads {
        query = query.Preload(preload) // Add each preload configuration
    }
    if err := query.Find(&results).Error; err != nil {
        return nil, err
    }
    return results, nil
}

// Update updates an entity by ID and checks ownership
func (s *GenericService[T]) UpdateService(c *gin.Context, id any, updatedEntity *T) (*T, error) {
	var entity T
	query := s.DB

	// Find the existing entity by ID
	if err := query.First(&entity, id).Error; err != nil {
		return nil, err
	}

	// Check ownership
	isOwner, err := s.CheckOwnership(c, entity)
	if err != nil || !isOwner {
		return nil, err
	}

	// Update the entity with the new values
	if err := s.DB.Model(&entity).Updates(updatedEntity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete deletes an entity by ID and checks ownership
func (s *GenericService[T]) DeleteService(c *gin.Context, id any) error {
	var entity T

	// Retrieve the record by ID
	if err := s.DB.First(&entity, id).Error; err != nil {
		return err // Record not found or another error
	}

	// Check ownership
	isOwner, err := s.CheckOwnership(c, entity)
	if err != nil || !isOwner {
		return err
	}

	// Delete the record
	if err := s.DB.Delete(&entity).Error; err != nil {
		return err // Error during deletion
	}

	return nil
}