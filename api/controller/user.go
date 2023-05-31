package controller

import (
	"ddd-go/application"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userApplication application.IUserApplication
}

func NewUserController(userApplication application.IUserApplication) *UserController {
	return &UserController{
		userApplication: userApplication,
	}
}

type UserResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (controller *UserController) GetUsers(c echo.Context) error {
	u, err := controller.userApplication.GetUsers(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UserResponse{
		ID:    uint64(u.ID),
		Email: string(u.Email),
		Name:  string(u.Name),
	})
}
