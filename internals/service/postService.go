package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/internals/domain"
)

type postService struct {
	postRepo domain.PostRepository
	userRepo domain.UserRepository
}

func NewPostService(postRepo domain.PostRepository, userRepo domain.UserRepository) domain.PostService {
	return &postService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (s *postService) CreatePost(ctx context.Context, req *domain.CreatePostRequest) (*sqlc.Post, error) {
	if req.UserID == "" {
		return nil, errors.New("user ID is required")
	}
	if req.Content == "" {
		return nil, errors.New("content is required")
	}

	// Parse and verify user exists
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Create post params
	params := sqlc.CreatePostParams{
		UserID:     userID,
		Content:    req.Content,
		ImageUrl:   req.ImageURL,
		LikesCount: 0,
	}

	return s.postRepo.CreatePost(ctx, params)
}

func (s *postService) GetPost(ctx context.Context, id string) (*sqlc.Post, error) {
	if id == "" {
		return nil, errors.New("post ID is required")
	}

	postID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid post ID format")
	}

	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	if post == nil {
		return nil, errors.New("post not found")
	}

	return post, nil
}

func (s *postService) DeletePost(ctx context.Context, id, userID string) error {
	if id == "" {
		return errors.New("post ID is required")
	}
	if userID == "" {
		return errors.New("user ID is required")
	}

	postID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid post ID format")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return fmt.Errorf("failed to get post: %w", err)
	}
	if post == nil {
		return errors.New("post not found")
	}

	// Check if user owns the post
	if post.UserID != userUUID {
		return errors.New("unauthorized: you can only delete your own posts")
	}

	return s.postRepo.DeletePost(ctx, postID)
}

func (s *postService) ListPosts(ctx context.Context, limit, offset int) ([]*sqlc.Post, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	params := sqlc.ListPostsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	return s.postRepo.ListPosts(ctx, params)
}
