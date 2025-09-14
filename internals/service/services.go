package service

import "github.com/jafferhussain11/celeb-social/internals/domain"

type Services struct {
	User       domain.UserService
	Post       domain.PostService
	Friendship domain.FriendshipService
}

func NewServices(repos *domain.Repositories) *Services {
	return &Services{
		User:       NewUserService(repos.User),
		Post:       NewPostService(repos.Post, repos.User),
		Friendship: NewFriendshipService(repos.Friendship, repos.User),
	}
}
