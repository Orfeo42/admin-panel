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
		echoCtx = utils.SetTitle(echoCtx, "Cliente")

		filter := repositories.NewCustomerFilter()

		utils.SetPageNumber(echoCtx, filter.Page)

		items, err := repositories.GetAllCustomerWithTotals(filter)
		if err != nil {
			return err
		}

		return utils.Render(pages.CustomerListView(*items, filter), echoCtx)
	})

	customerGroup.GET("/filter", func(echoCtx echo.Context) error {

		filter := GetCustomerFilterFromContext(echoCtx)

		utils.SetPageNumber(echoCtx, filter.Page)

		items, err := repositories.GetAllCustomerWithTotals(filter)
		if err != nil {
			return err
		}
		return utils.Render(pages.AllCustomerRowsShow(*items), echoCtx)
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

		filter := repositories.NewInvoiceFilter()

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

func GetCustomerFilterFromContext(echoCtx echo.Context) repositories.CustomerFilter {

	filter := repositories.NewCustomerFilter()

	page := utils.StringToInt(echoCtx.QueryParam("page"))
	if page != nil {
		filter.Page = *page
	}

	filter.Name = utils.StringToString(echoCtx.FormValue("name"))
	filter.TotalAmountFrom = utils.StringToAmount(echoCtx.FormValue("totalAmountFrom"))
	filter.TotalAmountTo = utils.StringToAmount(echoCtx.FormValue("totalAmountTo"))
	filter.TotalToPayFrom = utils.StringToAmount(echoCtx.FormValue("totalToPayFrom"))
	filter.TotalToPayTo = utils.StringToAmount(echoCtx.FormValue("totalToPayTo"))
	filter.IsPaid = isPaidToBool(echoCtx.FormValue("isPaid"))
	return filter

}
