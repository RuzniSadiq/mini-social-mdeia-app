package post_management

import (
	"mini-social-media/api/generics"
	"time"
)

// Post represents a social media post
type Post struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	UserID    int      `json:"user_id" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Foreign key to User table
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` // Cascade delete on Post deletion
	Likes    []Like    `json:"likes" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`    // Cascade delete on Post deletion
}

// Comment represents a comment on a post
type Comment struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Foreign keys
	PostID int  `json:"post_id" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` // References Post
	UserID int `json:"user_id" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // References User

}

// Like represents a like on a post or comment
type Like struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`

	// Foreign keys
	PostID *int `json:"post_id" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` 
	UserID int `json:"user_id" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

}


// PostService is responsible for handling requests related to posts
type PostService struct {
	*generics.GenericService[Post]
}

// CommentService is responsible for handling requests related to comments
type CommentService struct {
	*generics.GenericService[Comment]
}

// LikeService is responsible for handling business logic related to likes
type LikeService struct {
	*generics.GenericService[Like]
}


// PostController is responsible for handling requests related to posts
type PostController struct {
	*generics.GenericController[Post]
}

type CommentController struct {
	*generics.GenericController[Comment]
}

type LikeController struct {
	*generics.GenericController[Like]
}