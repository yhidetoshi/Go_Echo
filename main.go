package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./handler"
)

func main() {
	// Echoインスタンス作成
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())

	// サーバー起動
	e.Start(":1323") 
}
