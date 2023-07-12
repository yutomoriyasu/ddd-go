package main

import (
	"ddd-go/infrastructure/postgres"
	"ddd-go/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, dbClose, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer dbClose()

	router := router.NewRouter(e, db)
	router.Start()
}
