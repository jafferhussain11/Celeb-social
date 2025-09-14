package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/internals/domain"
)

type friendshipService struct {
	friendshipRepo domain.FriendshipRepository
	userRepo       domain.UserRepository
}

func NewFriendshipService(friendshipRepo domain.FriendshipRepository, userRepo domain.UserRepository) domain.FriendshipService {
	return &friendshipService{
		friendshipRepo: friendshipRepo,
		userRepo:       userRepo,
	}
}

func (s *friendshipService) SendFriendRequest(ctx context.Context, req *domain.FriendRequestRequest) (*sqlc.Friendship, error) {
	if req.RequesterID == "" {
		return nil, errors.New("requester ID is required")
	}
	if req.AddresseeID == "" {
		return nil, errors.New("addressee ID is required")
	}
	if req.RequesterID == req.AddresseeID {
		return nil, errors.New("cannot send friend request to yourself")
	}

	// Parse UUIDs
	requesterID, err := uuid.Parse(req.RequesterID)
	if err != nil {
		return nil, errors.New("invalid requester ID format")
	}

	addresseeID, err := uuid.Parse(req.AddresseeID)
	if err != nil {
		return nil, errors.New("invalid addressee ID format")
	}

	// Verify both users exist
	requester, err := s.userRepo.GetUserByID(ctx, requesterID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify requester: %w", err)
	}
	if requester == nil {
		return nil, errors.New("requester not found")
	}

	addressee, err := s.userRepo.GetUserByID(ctx, addresseeID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify addressee: %w", err)
	}
	if addressee == nil {
		return nil, errors.New("addressee not found")
	}

	// Check if friendship already exists
	existingParams := sqlc.GetFriendshipByUsersParams{
		RequesterID: requesterID,
		AddresseeID: addresseeID,
	}
	existingFriendship, err := s.friendshipRepo.GetFriendshipByUsers(ctx, existingParams)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing friendship: %w", err)
	}
	if existingFriendship != nil {
		return nil, errors.New("friendship request already exists")
	}

	// Create friendship
	params := sqlc.CreateFriendshipParams{
		RequesterID: requesterID,
		AddresseeID: addresseeID,
		Status:      "pending",
	}

	return s.friendshipRepo.CreateFriendship(ctx, params)
}

func (s *friendshipService) GetUserFriends(ctx context.Context, userID string) ([]*sqlc.Friendship, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	user, err := s.userRepo.GetUserByID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return s.friendshipRepo.GetUserFriends(ctx, userUUID)
}

func (s *friendshipService) RemoveFriend(ctx context.Context, id, userID string) error {
	if id == "" {
		return errors.New("friendship ID is required")
	}
	if userID == "" {
		return errors.New("user ID is required")
	}

	friendshipID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid friendship ID format")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	friendship, err := s.friendshipRepo.GetFriendshipByID(ctx, friendshipID)
	if err != nil {
		return fmt.Errorf("failed to get friendship: %w", err)
	}
	if friendship == nil {
		return errors.New("friendship not found")
	}

	// Check if user is part of this friendship
	if friendship.RequesterID != userUUID && friendship.AddresseeID != userUUID {
		return errors.New("unauthorized: you can only remove your own friendships")
	}

	return s.friendshipRepo.DeleteFriendship(ctx, friendshipID)
}
