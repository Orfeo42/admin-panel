package controllers

import (
	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/repositories"
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

		utils.SetPageNumber(echoCtx, 1)

		id, err := utils.StringToUint(stringId)
		if err != nil {
			return err
		}

		item, err := repositories.GetCustomerByID(*id)
		if err != nil {
			return err
		}

		echoCtx = utils.SetTitle(echoCtx, "Customer detail for customer: "+item.Name)

		filter := repositories.NewBaseFilter()

		filter.CustomerID = id

		utils.SetPageNumber(echoCtx, filter.Page)

		invoiceList, err := repositories.GetAllInvoice(filter)
		if err != nil {
			return err
		}

		return utils.Render(pages.CustomerView(*item, invoiceList, filter), echoCtx)
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
