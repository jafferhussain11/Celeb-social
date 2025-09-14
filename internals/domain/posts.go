package domain

import (
	"context"

	"github.com/google/uuid"
)

type CreatePostRequest struct {
	UserID   string  `json:"user_id" validate:"required"`
	Content  string  `json:"content" validate:"required"`
	ImageURL *string `json:"image_url"`
}

type PostService interface {
	CreatePost(ctx context.Context, req *CreatePostRequest) (*sqlc.Post, error)
	GetPost(ctx context.Context, id string) (*sqlc.Post, error)
	DeletePost(ctx context.Context, id, userID string) error
	ListPosts(ctx context.Context, limit, offset int) ([]*sqlc.Post, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, params sqlc.CreatePostParams) (*sqlc.Post, error)
	GetPostByID(ctx context.Context, id uuid.UUID) (*sqlc.Post, error)
	GetPostsByUserID(ctx context.Context, params sqlc.GetPostsByUserIDParams) ([]*sqlc.Post, error)
	ListPosts(ctx context.Context, params sqlc.ListPostsParams) ([]*sqlc.Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) error
}
