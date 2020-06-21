package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mori5602/statusAPI"
	"log"
	"path/filepath"
	"time"
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

	go func() {
		for {
			if err := j.Json.ReadFile(j.Path); err != nil {
				log.Println(err)
			}
			log.Println(j.Json)
			log.Println("AA")
			time.Sleep(2 * time.Minute)
			log.Println("BB")
		}

	}()

	e.GET("/status", func(c echo.Context) error {
		return j.Handler(c)
	})

	err := e.Start(":1234")
	if err != nil {
		log.Fatal(err)
	}
}
