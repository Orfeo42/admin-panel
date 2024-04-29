package handlers

import (
	"math/rand"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/order"
	"github.com/labstack/echo/v4"
)

func OrderListShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.OrderList)
	echoCtx = utils.SetTitle(echoCtx, "Orders")

	items := []data.Order{}

	for i := 0; i < 100; i++ {
		items = append(items, genRandomOrder())
	}

	return utils.Render(order.OrderListView(items), echoCtx)
}

func OrderShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.OrderAdd)
	echoCtx = utils.SetTitle(echoCtx, "Invoice")

	return utils.Render(order.OrderView(data.Order{}), echoCtx)
}

func genRandomOrder() data.Order {
	return data.Order{
		Customer: data.Customer{
			Name: utils.RandomString(25),
		},
		Amount: rand.Int(),
		Date:   utils.RandomDate(),
	}
}
