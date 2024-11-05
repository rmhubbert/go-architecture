package main

import "context"

type userService struct{}

func newUserService() *userService {
	return &userService{}
}

func (us *userService) findOne(ctx context.Context, id int) *user {
	return &user{
		id: id,
	}
}
