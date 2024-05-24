package main

import (
	"github.com/labstack/echo/v4/middleware"

	"github.com/Orfeo42/admin-panel/mvc/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	app := echo.New()

	app.HTTPErrorHandler = customHTTPErrorHandler
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} (${remote_ip}) ${latency_human}\n",
		Output: app.Logger.Output(),
	}))

	app.Static("/", "web")

	controllers.CustomerController(app)

	controllers.InvoiceController(app)

	controllers.HomeController(app)

	err := app.Start(":8080")
	if err != nil {
		log.Fatalf("Error Stating application: %+v", err)
	}
}

// TODO DA TESTARE E FARE PER BENE
func customHTTPErrorHandler(err error, echoCtx echo.Context) {
	/*code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	echoCtx.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := echoCtx.File(errorPage); err != nil {
		echoCtx.Logger().Error(err)
	}*/
}
