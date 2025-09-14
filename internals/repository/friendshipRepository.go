package repository

import "github.com/jafferhussain11/celeb-social/internals/domain"

type friendshipRepository struct {
	queries *sqlc.Queries
}

func NewFriendshipRepository(queries *sqlc.Queries) domain.FriendshipRepository {
	return &friendshipRepository{queries: queries}
}
