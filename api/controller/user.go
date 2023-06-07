package controller

import (
	"ddd-go/application"
	"ddd-go/controller/api"
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

func (controller *UserController) GetUsers(c echo.Context) error {
	users, err := controller.userApplication.GetUsers(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, api.NewGetUsersResponse(users))
}
