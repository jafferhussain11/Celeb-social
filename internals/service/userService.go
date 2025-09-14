package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/internals/domain"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*sqlc.User, error) {
	// Validation
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	if req.Username == "" {
		return nil, errors.New("username is required")
	}
	if req.FullName == "" {
		return nil, errors.New("full name is required")
	}

	// Check if user already exists
	existingUser, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	existingUser, err = s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing username: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// Create user params
	params := sqlc.CreateUserParams{
		Username:   strings.ToLower(req.Username),
		Email:      strings.ToLower(req.Email),
		FullName:   strings.TrimSpace(req.FullName),
		Bio:        req.Bio,
		IsVerified: false,
	}

	return s.userRepo.CreateUser(ctx, params)
}

func (s *userService) GetUser(ctx context.Context, id string) (*sqlc.User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, req *domain.UpdateUserRequest) (*sqlc.User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Get existing user
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Build update params
	params := sqlc.UpdateUserParams{
		ID:         userID,
		Username:   user.Username,   // Keep existing
		Email:      user.Email,      // Keep existing
		FullName:   user.FullName,   // Default to existing
		Bio:        user.Bio,        // Default to existing
		Avatar:     user.Avatar,     // Default to existing
		IsVerified: user.IsVerified, // Keep existing
	}

	// Update fields if provided
	if req.FullName != nil {
		params.FullName = strings.TrimSpace(*req.FullName)
	}
	if req.Bio != nil {
		params.Bio = req.Bio
	}
	if req.Avatar != nil {
		params.Avatar = req.Avatar
	}

	return s.userRepo.UpdateUser(ctx, params)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	return s.userRepo.DeleteUser(ctx, userID)
}

func (s *userService) ListUsers(ctx context.Context, limit, offset int) ([]*sqlc.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	params := sqlc.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	return s.userRepo.ListUsers(ctx, params)
}
