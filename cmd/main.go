package main

import (
	"github.com/Orfeo42/admin-panel/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/", "web")

	app.GET("/", handler.HomeShow)

	app.POST("/add-film/", handler.AddFilm)

	app.Start(":8080")
}
