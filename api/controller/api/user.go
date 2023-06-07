package api

import (
	"ddd-go/domain/model/user"
)

type User struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Users []User

type GetUsersResponse struct {
	Users Users `json:"users"`
}

// json
// {
// 	"users": [
// 		{
// 			"id": 1,
// 			"email": "test@example.com",
// 			"name": "test"
// 		}
// 	]
// }

func NewUser(u *user.User) User {
	return User{
		ID:    uint64(u.ID),
		Email: string(u.Email),
		Name:  string(u.Name),
	}
}

func NewUsers(users user.Users) Users {
	var us Users
	for _, u := range users {
		us = append(us, NewUser(u))
	}
	return us
}

func NewGetUsersResponse(users user.Users) GetUsersResponse {
	return GetUsersResponse{
		Users: NewUsers(users),
	}
}
