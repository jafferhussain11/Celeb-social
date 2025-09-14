package domain

import (
	"context"

	"github.com/google/uuid"
)

type FriendRequestRequest struct {
	RequesterID string `json:"requester_id" validate:"required"`
	AddresseeID string `json:"addressee_id" validate:"required"`
}

type FriendshipService interface {
	SendFriendRequest(ctx context.Context, req *FriendRequestRequest) (*sqlc.Friendship, error)
	GetUserFriends(ctx context.Context, userID string) ([]*sqlc.Friendship, error)
	RemoveFriend(ctx context.Context, id, userID string) error
}

type FriendshipRepository interface {
	CreateFriendship(ctx context.Context, params sqlc.CreateFriendshipParams) (*sqlc.Friendship, error)
	GetFriendshipByID(ctx context.Context, id uuid.UUID) (*sqlc.Friendship, error)
	GetFriendshipByUsers(ctx context.Context, params sqlc.GetFriendshipByUsersParams) (*sqlc.Friendship, error)
	GetUserFriends(ctx context.Context, userID uuid.UUID) ([]*sqlc.Friendship, error)
	UpdateFriendshipStatus(ctx context.Context, params sqlc.UpdateFriendshipStatusParams) (*sqlc.Friendship, error)
	DeleteFriendship(ctx context.Context, id uuid.UUID) error
}
