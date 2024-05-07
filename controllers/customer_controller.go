package controllers

import (
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/customer"
	"github.com/labstack/echo/v4"
)

func CustomerController(application *echo.Echo) {

	customerGroup := application.Group("/customer")

	customerGroup.GET("/list", func(echoCtx echo.Context) error {

		echoCtx = utils.SetPage(echoCtx, pages.CustomerList)
		echoCtx = utils.SetTitle(echoCtx, "Invoices")
		items, err := repositories.GetAllCustomer()
		if err != nil {
			return err
		}

		return utils.Render(customer.CustomerListView(items), echoCtx)
	})

	customerGroup.GET("/add", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, pages.CustomerAdd)
		echoCtx = utils.SetTitle(echoCtx, "Invoice")
		return utils.Render(customer.CustomerView(repositories.Customer{}), echoCtx)
	})

	customerGroup.POST("", func(echoCtx echo.Context) error {
		input := repositories.Customer{
			Name:    echoCtx.FormValue("name"),
			Surname: echoCtx.FormValue("surname"),
			Address: echoCtx.FormValue("address"),
			Email:   echoCtx.FormValue("email"),
			Phone:   echoCtx.FormValue("phone"),
		}
		result, err := repositories.CreateCustomer(input)
		if err != nil {
			return err
		}

		return utils.Render(customer.CustomerForm(result), echoCtx)
	})

}
