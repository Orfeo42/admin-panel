package main

import (
	"github.com/Orfeo42/admin-panel/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/", "web")

	app.GET("/", handler.HomeShow)

	invoiceGroup := app.Group("/invoice")

	invoiceGroup.GET("/list", handler.InvoiceListShow)

	invoiceGroup.GET("/add", handler.HomeShow)

	customerGroup := app.Group("/customer")

	customerGroup.GET("/list", handler.CustomerListShow)

	customerGroup.GET("/add", handler.HomeShow)

	orderGroup := app.Group("/order")

	orderGroup.GET("/list", handler.OrderListShow)

	orderGroup.GET("/add", handler.HomeShow)

	app.GET("/films", handler.FilmShow)

	app.POST("/add-film/", handler.AddFilm)

	app.Start(":8080")
}
