package post_management

import (
	"mini-social-media/api/generics"
	"mini-social-media/initializers"

	"github.com/gin-gonic/gin"
)


func NewCommentService() *CommentService {
	return &CommentService{
		GenericService: generics.NewGenericService[Comment](initializers.DB),
	}
}


// CreateComment creates a new comment using the generic CreateService
func (s *CommentService) Create(c *gin.Context, comment *Comment, user_id int) (*Comment, error) {
	comment.UserID = user_id
	return s.GenericService.CreateService(comment)
}

// GetCommentByID retrieves a comment by ID using the generic RetrieveService
func (s *CommentService) Retrieve(c *gin.Context, id any) (*Comment, error) {
	return s.GenericService.RetrieveService(id)
}

// GetAllComments retrieves all comments using the generic ListService
func (s *CommentService) List(c *gin.Context) ([]Comment, error) {
	return s.GenericService.ListService("posts")
}

// UpdateComment updates an existing comment and checks ownership
func (s *CommentService) Update(c *gin.Context, id any, comment *Comment) (*Comment, error) {
	return s.GenericService.UpdateService(c, id, comment)
}

// DeleteComment deletes a comment and checks ownership before deletion
func (s *CommentService) Delete(c *gin.Context, id any) error {
	return s.GenericService.DeleteService(c, id)
}