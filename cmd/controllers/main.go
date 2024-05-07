package main

import (
	"github.com/Orfeo42/admin-panel/controllers"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/home"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	app := echo.New()

	app.Static("/", "web")

	controllers.CustomerController(app)

	controllers.InvoiceController(app)

	controllers.OrderController(app)

	app.GET("/", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, pages.Home)
		echoCtx = utils.SetTitle(echoCtx, "Home Page")
		return utils.Render(home.HomeView(), echoCtx)
	})

	err := app.Start(":8080")
	if err != nil {
		log.Fatalf("Error Stating application: %+v", err)
	}
}
