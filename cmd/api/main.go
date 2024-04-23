package main

import (
	"github.com/Orfeo42/admin-panel/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/", "web")

	/*----------
	  Customer Group
	------------*/

	customerGroup := app.Group("/customer")

	customerGroup.GET("/list", handler.CustomerListShow)

	customerGroup.GET("/add", handler.CustomerShow)

	/*----------
	  Invoice Group
	------------*/

	invoiceGroup := app.Group("/invoice")

	invoiceGroup.GET("/list", handler.InvoiceListShow)

	invoiceGroup.GET("/add", handler.InvoiceShow)

	/*----------
	  Order Group
	------------*/

	orderGroup := app.Group("/order")

	orderGroup.GET("/list", handler.OrderListShow)

	orderGroup.GET("/add", handler.OrderShow)

	/*----------
	  Other
	------------*/

	app.GET("/", handler.HomeShow)

	app.GET("/films", handler.FilmShow)

	app.POST("/add-film/", handler.AddFilm)

	app.Start(":8080")
}
