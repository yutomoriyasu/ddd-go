package router

import (
	"ddd-go/application"
	"ddd-go/controller"
	"ddd-go/repository"

	"github.com/labstack/echo/v4"
)

type Router struct {
	e *echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	e.GET("/", hello)

	ur := repository.NewUserRepository()
	ua := application.NewUser(ur)
	uc := controller.NewUserController(ua)
	e.GET("/users", uc.GetUsers)

	return &Router{e: e}
}

func (r *Router) Start() {
	r.e.Logger.Fatal(r.e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
