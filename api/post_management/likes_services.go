package post_management

import (
	"mini-social-media/api/generics"
	"mini-social-media/initializers"

	"github.com/gin-gonic/gin"
)



func NewLikeService() *LikeService {
	return &LikeService{
		GenericService: generics.NewGenericService[Like](initializers.DB),
	}
}


// CreateLikes creates a new comment using the generic CreateService
func (s *LikeService) Create(c *gin.Context, like *Like, user_id int) (*Like, error) {
	like.UserID = user_id
	return s.GenericService.CreateService(like)
}

// GetLikeByID retrieves a like by ID using the generic RetrieveService
func (s *LikeService) Retrieve(c *gin.Context, id any) (*Like, error) {
	return s.GenericService.RetrieveService(id)
}

// GetAllLikes retrieves all likes using the generic ListService
func (s *LikeService) List(c *gin.Context) ([]Like, error) {
	return s.GenericService.ListService()
}

// UpdateLike updates an existing like and checks ownership
func (s *LikeService) Update(c *gin.Context, id any, like *Like) (*Like, error) {
	return s.GenericService.UpdateService(c, id, like)
}

// DeleteLike deletes a like and checks ownership before deletion
func (s *LikeService) Delete(c *gin.Context, id any) error {
	return s.GenericService.DeleteService(c, id)
}