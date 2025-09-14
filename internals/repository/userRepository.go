package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/internals/database/sqlc"
	"github.com/jafferhussain11/celeb-social/internals/domain"
)

type userRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) domain.UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*sqlc.User, error) {
	user, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*sqlc.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*sqlc.User, error) {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*sqlc.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, params sqlc.UpdateUserParams) (*sqlc.User, error) {
	user, err := r.queries.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *userRepository) ListUsers(ctx context.Context, params sqlc.ListUsersParams) ([]*sqlc.User, error) {
	users, err := r.queries.ListUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	// Convert []sqlc.User to []*sqlc.User
	result := make([]*sqlc.User, len(users))
	for i := range users {
		result[i] = &users[i]
	}

	return result, nil
}
