package repository

import "github.com/jafferhussain11/celeb-social/internals/domain"

type postRepository struct {
	queries *sqlc.Queries
}

func NewPostRepository(queries *sqlc.Queries) domain.PostRepository {
	return &postRepository{queries: queries}
}
