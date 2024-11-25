package generics

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



// NewGenericController returns a new instance of GenericController
func NewGenericController[T any](method GenericMain[T]) *GenericController[T] {
	return &GenericController[T]{
		Method: method,
	}
}

// Create handles the creation of a new entity
func (c *GenericController[T]) Create(ctx *gin.Context) {
	var req T
	fmt.Println("is this being triggered")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		CustomJsonResponseError(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	userID, exists := ctx.Get("user_id")

	if !exists {
		CustomJsonResponseError(ctx, http.StatusUnauthorized, "User not authenticated")
		return
	}


	// Ensure userID is of the correct type
	user_int, ok := userID.(int)
	if !ok {
		CustomJsonResponseError(ctx, http.StatusInternalServerError, "Invalid user ID format")
		return
	}
	

	createdEntity, err := c.Method.Create(ctx, &req, user_int)
	if err != nil {
		CustomJsonResponseError(ctx, http.StatusInternalServerError, "Error creating entity")
		return
	}

	CustomJsonResponseData(ctx, http.StatusCreated, createdEntity)
}

// Retrieve handles retrieving an entity by ID
func (c *GenericController[T]) Retrieve(ctx *gin.Context) {
	id := ctx.Param("id")
	entity, err := c.Method.Retrieve(ctx, id)
	if err != nil {
		CustomJsonResponseError(ctx, http.StatusNotFound, "Entity not found")
		return
	}

	CustomJsonResponseData(ctx, http.StatusOK, entity)
}

// List handles retrieving a list of entities
func (c *GenericController[T]) List(ctx *gin.Context) {
	entities, err := c.Method.List(ctx)
	if err != nil {
		CustomJsonResponseError(ctx, http.StatusInternalServerError, "Error retrieving entities")
		return
	}

	CustomJsonResponseData(ctx, http.StatusOK, entities)
}

// // Update handles updating an entity by ID
func (c *GenericController[T]) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req T
	if err := ctx.ShouldBindJSON(&req); err != nil {
		CustomJsonResponseError(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	updatedEntity, err := c.Method.Update(ctx, id, &req)
	if err != nil {
		CustomJsonResponseError(ctx, http.StatusForbidden, "Unauthorized or error updating entity")
		return
	}

	CustomJsonResponseData(ctx, http.StatusOK, updatedEntity)
}

// Delete handles deleting an entity by ID
func (c *GenericController[T]) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.Method.Delete(ctx, id)
	if err != nil {
		CustomJsonResponseError(ctx, http.StatusForbidden, "Unauthorized or error deleting entity")
		return
	}

	CustomJsonResponseData(ctx, http.StatusOK, "Entity deleted successfully")
}
