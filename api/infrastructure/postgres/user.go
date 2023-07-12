package postgres

import (
	"context"
	"ddd-go/domain/model/user"

	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() user.IRepository {
	return &userRepository{}
}

type userDTO struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

type userDTOs []*userDTO

func newUserDTO(u *user.User) *userDTO {
	return &userDTO{
		Name:  string(u.Name),
		Email: string(u.Email),
	}
}

func newUserDTOs(us user.Users) userDTOs {
	dto := make(userDTOs, len(us))
	for i, u := range us {
		dto[i] = newUserDTO(u)
	}
	return dto
}

func (u *userDTO) toUser() (*user.User, error) {
	id := user.NewID(uint64(u.ID))
	email, err := user.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}
	name, err := user.NewName(u.Name)
	if err != nil {
		return nil, err
	}
	usr := user.NewUser(id, email, name)
	return usr, nil
}

func (us userDTOs) toUsers() (user.Users, error) {
	users := make(user.Users, len(us))
	for i, u := range us {
		usr, err := u.toUser()
		if err != nil {
			return nil, err
		}
		users[i] = usr
	}
	return users, nil
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
