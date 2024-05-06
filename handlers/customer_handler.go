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
	items, err := data.GetAllCustomer()
	if err != nil {
		return err
	}

	return utils.Render(customer.CustomerListView(items), echoCtx)
}

func CustomerShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.CustomerAdd)
	echoCtx = utils.SetTitle(echoCtx, "Invoice")
	return utils.Render(customer.CustomerView(data.Customer{}), echoCtx)
}

func CustomerAdd(echoCtx echo.Context) error {
	input := data.Customer{
		Name:    echoCtx.FormValue("name"),
		Surname: echoCtx.FormValue("surname"),
		Address: echoCtx.FormValue("address"),
		Email:   echoCtx.FormValue("email"),
		Phone:   echoCtx.FormValue("phone"),
	}
	result, err := data.CreateCustomer(input)
	if err != nil {
		return err
	}

	return utils.Render(customer.CustomerForm(result), echoCtx)
}

func ShowModal(echoCtx echo.Context) error {
	return utils.Render(customer.CustomerModal(), echoCtx)
}
