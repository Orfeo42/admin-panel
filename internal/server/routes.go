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

	customers.RegisterRoutes(e)

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

	invoiceGroup.POST("", controller.CreateHandler)

	invoiceGroup.GET("/add", controller.CreatePageHandler)

	invoiceGroup.GET("/list", controller.ReadAllHandler)

	invoiceGroup.GET("/filter", controller.FilterHandler)

	invoiceGroup.PUT("/:id", controller.UpdateHandler)

	invoiceGroup.PUT("/:id/pay", controller.PayByID)

	invoiceGroup.GET("/:id/edit", controller.UpdatePageHandler)

	invoiceGroup.GET("/:id", controller.ReadHandler)

}
