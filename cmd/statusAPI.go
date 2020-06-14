package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mori5602/statusAPI"
	"log"
	"path/filepath"
)

func main() {
	//	Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Static("static"))

	// JSON
	j := statusAPI.NewStatusFactory(filepath.Join("testdata", "ajax.json"))
	e.GET("/status", func(c echo.Context) error {
		return j.Handler(c)
	})

	err := e.Start(":1234")
	if err != nil {
		log.Fatal(err)
	}
}
