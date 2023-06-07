package repository

import (
	"context"
	"ddd-go/domain/model/user"
)

type userRepository struct{}

func NewUserRepository() user.IRepository {
	return &userRepository{}
}

func (u *userRepository) FindAll(ctx context.Context) (user.Users, error) {
	id := user.NewID(1)
	email, err := user.NewEmail("test@example.com")
	if err != nil {
		return nil, err
	}
	name, err := user.NewName("test")
	if err != nil {
		return nil, err
	}
	usr := user.NewUser(id, email, name)
	return user.Users{usr}, nil
}
