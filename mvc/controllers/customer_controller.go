package controllers

import (
	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/pages"
	"github.com/labstack/echo/v4"
)

func CustomerController(application *echo.Echo) {

	customerGroup := application.Group("/customer")

	customerGroup.GET("/list", func(echoCtx echo.Context) error {

		echoCtx = utils.SetPage(echoCtx, enum.CustomerList)
		echoCtx = utils.SetTitle(echoCtx, "Customer")
		items, err := repositories.GetAllCustomerWithTotals()
		if err != nil {
			return err
		}

		return utils.Render(pages.CustomerListView(*items), echoCtx)
	})

	customerGroup.GET("/:id/info", func(echoCtx echo.Context) error {
		stringId := echoCtx.Param("id")

		echoCtx = utils.SetPage(echoCtx, enum.CustomerList)

		id, err := utils.StringToUint(stringId)
		if err != nil {
			return err
		}

		item, err := repositories.GetCustomerByID(*id)
		if err != nil {
			return err
		}

		echoCtx = utils.SetTitle(echoCtx, "Customer detail for customer: "+item.Name)

		invoiceList, err := repositories.GetAllInvoiceByCustomerID(*id, nil)
		if err != nil {
			return err
		}

		return utils.Render(pages.CustomerView(*item, *invoiceList), echoCtx)
	})

	customerGroup.GET("/search", func(echoCtx echo.Context) error {
		name := echoCtx.QueryParam("name")

		customerList, err := repositories.SearchCustomerByName(name)
		if err != nil {
			return err
		}

		return utils.Render(pages.CustomerSearchView(*customerList), echoCtx)
	})
}
