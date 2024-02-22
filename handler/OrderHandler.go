package handler

import (
	"math/rand"

	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/model"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/order"
	"github.com/labstack/echo/v4"
)

func OrderListShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.OrderList)
	echoCtx = utils.SetTitle(echoCtx, "Orders")

	items := []model.OrderModel{}

	for i := 0; i < 100; i++ {
		items = append(items, genRandomOrder())
	}

	return utils.Render(order.OrderView(items), echoCtx)
}

func genRandomOrder() model.OrderModel {
	return model.OrderModel{
		Customer: utils.RandomString(25),
		Amount:   rand.Float64(),
		Date:     utils.RandomDate(),
	}
}
