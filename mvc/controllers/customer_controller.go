package controllers

import (
	"fmt"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/pages/viewcustomer"
	"github.com/labstack/echo/v4"
)

func getCustomerFilterFromContext(echoCtx echo.Context) repositories.CustomerFilter {

	filter := repositories.NewCustomerFilter()

	page, err := utils.StringToInt(echoCtx.QueryParam("page"))
	if err == nil {
		filter.Page = *page
	}

	filter.Name = utils.ParseString(echoCtx.FormValue("name"))
	filter.TotalAmountFrom = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalAmountFrom"))
	filter.TotalAmountTo = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalAmountTo"))
	filter.TotalToPayFrom = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalToPayFrom"))
	filter.TotalToPayTo = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalToPayTo"))
	filter.IsPaid = isPaidToBool(echoCtx.FormValue("isPaid"))
	return filter

}

func CustomerController(application *echo.Echo) {

	customerGroup := application.Group("/customer")

	customerGroup.GET("/list", func(echoCtx echo.Context) error {

		filter := getCustomerFilterFromContext(echoCtx)

		items, err := repositories.GetAllCustomerWithTotals(filter)
		if err != nil {
			return err
		}

		utils.SetPageNumber(echoCtx, filter.Page)

		utils.SetPage(echoCtx, enum.CustomerList)

		utils.SetTitle(echoCtx, "Cliente")

		customerListParams := viewcustomer.CustomerListParams{
			Items:  *items,
			Filter: filter,
		}

		return utils.Render(viewcustomer.CustomerListView(customerListParams), echoCtx)
	})

	customerGroup.GET("/:id/info", func(echoCtx echo.Context) error {

		stringId := echoCtx.Param("id")

		id, err := utils.StringToUintPtr(stringId)
		if err != nil {
			return err
		}

		customer, err := repositories.GetCustomerByID(*id)
		if err != nil {
			return err
		}

		filter := repositories.NewInvoiceFilter()
		filter.CustomerID = &customer.ID

		invoiceList, err := repositories.GetAllInvoice(filter)
		if err != nil {
			return err
		}

		utils.SetPageNumber(echoCtx, filter.Page)

		utils.SetPage(echoCtx, enum.CustomerList)

		utils.SetTitle(echoCtx, fmt.Sprintf("%s - customer detail", customer.Name))

		customerDetailParams := viewcustomer.CustomerDetailParams{
			Item:        *customer,
			InvoiceList: invoiceList,
			Filter:      filter,
		}

		return utils.Render(viewcustomer.CustomerDetail(customerDetailParams), echoCtx)
	})

	customerGroup.GET("/filter", func(echoCtx echo.Context) error {

		filter := getCustomerFilterFromContext(echoCtx)

		items, err := repositories.GetAllCustomerWithTotals(filter)
		if err != nil {
			return err
		}

		utils.SetPageNumber(echoCtx, filter.Page)

		return utils.Render(viewcustomer.AllCustomerRowsShow(*items), echoCtx)
	})

	customerGroup.GET("/search", func(echoCtx echo.Context) error {
		name := echoCtx.QueryParam("name")

		customerList, err := repositories.SearchCustomerByName(name)
		if err != nil {
			return err
		}

		return utils.Render(viewcustomer.CustomerSearch(*customerList), echoCtx)
	})
}
