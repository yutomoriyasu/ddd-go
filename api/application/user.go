package application

import (
	"context"
	"ddd-go/domain/model/user"
)

type IUserApplication interface {
	GetUsers(ctx context.Context) (user.Users, error)
}

type userApplication struct {
	userRepo user.IRepository
}

func NewUser(userRepo user.IRepository) IUserApplication {
	return &userApplication{
		userRepo: userRepo,
	}
}

func (u *userApplication) GetUsers(ctx context.Context) (user.Users, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
