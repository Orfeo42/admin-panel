package main

import (
	"github.com/Orfeo42/admin-panel/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

	customerGroup.GET("/modal", handlers.ShowModal)

	/*----------
	  Invoice Group
	------------*/

	invoiceGroup := app.Group("/invoice")

	invoiceGroup.GET("/list", handlers.InvoiceListShow)

	invoiceGroup.GET("/filter", handlers.InvoiceFilter)

	invoiceGroup.GET("/add", handlers.InvoiceShow)

	customerGroup.POST("", handlers.InvoiceAdd)

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

	err := app.Start(":8080")
	if err != nil {
		log.Fatalf("Error Stating application: %+v", err)
	}
}
