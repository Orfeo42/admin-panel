package controllers

import (
	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/view/pages"
	"math/rand"

	"github.com/Orfeo42/admin-panel/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/labstack/echo/v4"
)

func OrderController(application *echo.Echo) {
	orderGroup := application.Group("/order")

	orderGroup.GET("/list", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, enum.OrderList)
		echoCtx = utils.SetTitle(echoCtx, "Orders")

		var items []repositories.Order

		for i := 0; i < 100; i++ {
			items = append(items, genRandomOrder())
		}

		return utils.Render(pages.OrderListView(items), echoCtx)
	})

	orderGroup.GET("/add", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, enum.OrderAdd)
		echoCtx = utils.SetTitle(echoCtx, "Invoice")

		return utils.Render(pages.OrderView(repositories.Order{}), echoCtx)
	})
}

func genRandomOrder() repositories.Order {
	return repositories.Order{
		Customer: repositories.Customer{
			Name: utils.RandomString(25),
		},
		Amount: rand.Int(),
		Date:   utils.RandomDate(),
	}
}
