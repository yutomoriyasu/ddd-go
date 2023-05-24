package main

import (
	"ddd-go/domain/model/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", hello) // ローカル環境の場合、http://localhost:8080/ にGETアクセスされるとhelloハンドラーを実行する

	e.GET("/users", getUsers)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type UserResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func getUsers(c echo.Context) error {
	id := user.NewID(1)
	email, err := user.NewEmail("test@example.com")
	if err != nil {
		return err
	}
	name, err := user.NewName("test")
	if err != nil {
		return err
	}
	u := user.NewUser(id, email, name)

	return c.JSON(http.StatusOK, UserResponse{
		ID:    uint64(u.ID),
		Email: string(u.Email),
		Name:  string(u.Name),
	})
}
