package handlers

import (
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/customer"
	"github.com/labstack/echo/v4"
)

func CustomerListShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.CustomerList)
	echoCtx = utils.SetTitle(echoCtx, "Invoices")

	items := []data.CustomerModel{}

	for i := 0; i < 100; i++ {
		items = append(items, genRandomCustomer())
	}

	return utils.Render(customer.CustomerListView(items), echoCtx)
}

func CustomerShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.CustomerAdd)
	echoCtx = utils.SetTitle(echoCtx, "Invoice")

	return utils.Render(customer.CustomerView(data.CustomerModel{}), echoCtx)
}

func genRandomCustomer() data.CustomerModel {
	return data.CustomerModel{
		Name:    utils.RandomString(25),
		Surname: utils.RandomString(25),
	}
}
