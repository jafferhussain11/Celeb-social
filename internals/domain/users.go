package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/internals/database/sqlc"
)

type CreateUserRequest struct {
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	FullName string  `json:"full_name" validate:"required"`
	Bio      *string `json:"bio"`
	//add password field
}

type UpdateUserRequest struct {
	FullName *string `json:"full_name"`
	Bio      *string `json:"bio"`
	Avatar   *string `json:"avatar"`
}

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*sqlc.User, error)
	GetUser(ctx context.Context, id string) (*sqlc.User, error)
	UpdateUser(ctx context.Context, id string, req *UpdateUserRequest) (*sqlc.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, limit, offset int) ([]*sqlc.User, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*sqlc.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*sqlc.User, error)
	GetUserByUsername(ctx context.Context, username string) (*sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (*sqlc.User, error)
	UpdateUser(ctx context.Context, params sqlc.UpdateUserParams) (*sqlc.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ListUsers(ctx context.Context, params sqlc.ListUsersParams) ([]*sqlc.User, error)
}
