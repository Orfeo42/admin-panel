package main

import (
	"github.com/Orfeo42/admin-panel/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	app.Static("/", "web")

	/*----------
	  Customer Group
	------------*/

	customerGroup := app.Group("/customer")

	customerGroup.GET("/list", handlers.CustomerListShow)

	customerGroup.GET("/add", handlers.CustomerShow)

	customerGroup.POST("", handlers.CustomerAdd)

	/*----------
	  Invoice Group
	------------*/

	invoiceGroup := app.Group("/invoice")

	invoiceGroup.GET("/list", handlers.InvoiceListShow)

	invoiceGroup.GET("/add", handlers.InvoiceShow)

	/*----------
	  Order Group
	------------*/

	orderGroup := app.Group("/order")

	orderGroup.GET("/list", handlers.OrderListShow)

	orderGroup.GET("/add", handlers.OrderShow)

	/*----------
	  Other
	------------*/

	app.GET("/", handlers.HomeShow)

	app.Start(":8080")
}
