package main

import (
	"github.com/Orfeo42/admin-panel/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/", handler.HomeShow)

	app.Start(":8080")
}
