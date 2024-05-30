package server

import (
	"net/http"

	"admin-panel/cmd/web"
	"admin-panel/cmd/web/customers"
	"admin-panel/cmd/web/home"
	"admin-panel/cmd/web/invoices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))

	fileServer := http.FileServer(http.FS(web.Files))

	e.GET("/assets/*", echo.WrapHandler(fileServer)) //TODO

	registerHomeRoutes(e)

	registerCustomerRoutes(e)

	registerInvoiceRoutes(e)

	return e
}

func registerHomeRoutes(application *echo.Echo) {

	homeGroup := application.Group("")

	controller := home.HomeControllerInstance()

	homeGroup.GET("", controller.HomeHandler)
}

func registerInvoiceRoutes(application *echo.Echo) {
	invoiceGroup := application.Group("/invoice")
	controller := invoices.InvoiceControllerInstance()

	invoiceGroup.GET("", controller.ReadAllPageHandler)
	invoiceGroup.GET("/:id", controller.ReadPageHandler)
	invoiceGroup.GET("/filter", controller.FilterHandler)

	invoiceGroup.GET("/add", controller.CreatePageHandler)
	invoiceGroup.POST("", controller.CreateHandler)

	invoiceGroup.GET("/:id/edit", controller.UpdatePageHandler)
	invoiceGroup.PUT("/:id", controller.UpdateHandler)
	invoiceGroup.PUT("/:id/pay", controller.PayByID)
}

func registerCustomerRoutes(application *echo.Echo) {

	customerGroup := application.Group("/customer")
	controller := customers.CustomerControllerInstance()

	customerGroup.GET("", controller.ReadAllPageHandler)
	customerGroup.GET("/:id", controller.ReadPageHandler)
	customerGroup.GET("/filter", controller.FilterHandler)
	customerGroup.GET("/search", controller.SearchByNameHandler)

	customerGroup.GET("/add", controller.CreatePageHandler)
	customerGroup.POST("", controller.CreateHandler)

	customerGroup.GET("/:id/edit", controller.UpdatePageHandler)
	customerGroup.PUT("/:id", controller.UpdateHandler)
}
