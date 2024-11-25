package generics

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Service defines the basic CRUD operations for any entity
type GenericMain[T any] interface {
	Create(ctx *gin.Context, entity *T, user_id int) (*T, error)
	Retrieve(ctx *gin.Context, id any) (*T, error)
	List(ctx *gin.Context) ([]T, error)
	Update(ctx *gin.Context, id any, updatedEntity *T) (*T, error)
	Delete(ctx *gin.Context, id any) error
}

// GenericController defines common CRUD operations for any resource
type GenericController[T any] struct {
	Method GenericMain[T]
}

// GenericService encapsulates CRUD operations for any entity type T
type GenericService[T any] struct {
	DB *gorm.DB
}