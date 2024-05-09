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
		echoCtx = utils.SetTitle(echoCtx, "Customer")
		items, err := repositories.GetAllCustomerWithTotals()
		if err != nil {
			return err
		}

		return utils.Render(customer.CustomerListView(*items), echoCtx)
	})

	customerGroup.GET("/:id/info", func(echoCtx echo.Context) error {
		id := echoCtx.Param("id")
		echoCtx = utils.SetPage(echoCtx, pages.CustomerList)

		item, err := repositories.GetCustomerByIDString(id)
		if err != nil {
			return err
		}
		echoCtx = utils.SetTitle(echoCtx, "Customer detail for customer: "+item.Name)
		invoiceList, err := repositories.GetAllInvoiceByCustomerID(id, nil)
		if err != nil {
			return err
		}

		return utils.Render(customer.CustomerView(*item, *invoiceList), echoCtx)
	})

	/*customerGroup.POST("", func(echoCtx echo.Context) error {
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

		return utils.Render(customer.CustomerData(result), echoCtx)
	})*/

	customerGroup.GET("/search", func(echoCtx echo.Context) error {
		name := echoCtx.QueryParam("name")

		customerList, err := repositories.SearchCustomerByName(name)
		if err != nil {
			return err
		}

		return utils.Render(customer.CustomerSearchView(*customerList), echoCtx)
	})
}
