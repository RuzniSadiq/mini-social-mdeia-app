package post_management

import "mini-social-media/api/generics"



// NewPostController returns a new instance of PostController
func NewPostController() *PostController {
	postService := NewPostService()
	return &PostController{
		GenericController: generics.NewGenericController[Post](postService),
	}
}

// NewCommentController returns a new instance of CommentController
func NewCommentController() *CommentController {
	commentService := NewCommentService()
	return &CommentController{
		GenericController: generics.NewGenericController[Comment](commentService),
	}
}

// NewLikeController returns a new instance of LikeController
func NewLikeController() *LikeController {
	likeService := NewLikeService()
	return &LikeController{
		GenericController: generics.NewGenericController[Like](likeService),
	}
}

