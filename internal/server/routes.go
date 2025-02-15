package server

import (
	"admin-panel/cmd/web/update"
	"net/http"

	"admin-panel/cmd/web"
	"admin-panel/cmd/web/charts"
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

	e.GET("/assets/*", echo.WrapHandler(fileServer))

	update.RegisterRoutes(e)

	home.RegisterRoutes(e)

	charts.RegisterRoutes(e)

	customers.RegisterCustomerRoutes(e)

	invoices.RegisterRoutes(e)

	return e
}
