package post_management

import (
	"mini-social-media/api/generics"
	"mini-social-media/initializers"

	"github.com/gin-gonic/gin"
)



func NewPostService() *PostService {
	return &PostService{
		GenericService: generics.NewGenericService[Post](initializers.DB),
	}
}

// CreatePost creates a new post using the generic CreateService
func (s *PostService) Create(c *gin.Context, post *Post, user_id int) (*Post, error) {
	post.UserID = user_id
	return s.GenericService.CreateService(post)
}

// GetPostByID retrieves a post by ID using the generic RetrieveService
func (s *PostService) Retrieve(c *gin.Context, id any) (*Post, error) {
	return s.GenericService.RetrieveService(id)
}

// GetAllPosts retrieves all posts using the generic ListService
func (s *PostService) List(c *gin.Context) ([]Post, error) {
	return s.GenericService.ListService("Comments", "Likes")
}

// UpdatePost updates an existing post and checks ownership
func (s *PostService) Update(c *gin.Context, id any, post *Post) (*Post, error) {
	return s.GenericService.UpdateService(c, id, post)
}

// DeletePost deletes a post and checks ownership before deletion
func (s *PostService) Delete(c *gin.Context, id any) error {
	return s.GenericService.DeleteService(c, id)
}