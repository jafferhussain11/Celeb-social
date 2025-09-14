package repository

import (
	"database/sql"

	"github.com/jafferhussain11/celeb-social/internals/domain"
)

type Repositories struct {
	User       domain.UserRepository
	Post       domain.PostRepository
	Friendship domain.FriendshipRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:       NewUserRepository(db),
		Post:       NewPostRepository(db),
		Friendship: NewFriendshipRepository(db),
	}
}
