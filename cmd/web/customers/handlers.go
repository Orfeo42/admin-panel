package customers

import (
	"fmt"

	"admin-panel/cmd/enum"
	"admin-panel/internal/database"
	"admin-panel/utils"

	"github.com/labstack/echo/v4"
)

func getCustomerFilterFromContext(echoCtx echo.Context) database.CustomerFilter {

	filter := database.NewCustomerFilter()

	page, err := utils.StringToInt(echoCtx.QueryParam("page"))
	if err == nil {
		filter.Page = *page
	}

	filter.Name = utils.ParseString(echoCtx.FormValue("name"))
	filter.TotalAmountFrom = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalAmountFrom"))
	filter.TotalAmountTo = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalAmountTo"))
	filter.TotalToPayFrom = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalToPayFrom"))
	filter.TotalToPayTo = utils.StringToAmountPtrNoErr(echoCtx.FormValue("totalToPayTo"))
	//filter.IsPaid = isPaidToBool(echoCtx.FormValue("isPaid"))
	return filter

}

func RegisterRoutes(application *echo.Echo) {

	customerGroup := application.Group("/customer")

	customerGroup.GET("/list", listHandler)

	customerGroup.GET("/:id/info", infoHandler)

	customerGroup.GET("/filter", filterHandler)

	customerGroup.GET("/search", searchHandler)
}

func listHandler(echoCtx echo.Context) error {

	filter := getCustomerFilterFromContext(echoCtx)

	items, err := database.GetAllCustomerWithTotals(filter)
	if err != nil {
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	utils.SetPage(echoCtx, enum.CustomerList)

	utils.SetTitle(echoCtx, "Cliente")

	customerListParams := CustomerListParams{
		Items:  *items,
		Filter: filter,
	}

	return utils.Render(CustomerListView(customerListParams), echoCtx)
}

func infoHandler(echoCtx echo.Context) error {

	stringId := echoCtx.Param("id")

	id, err := utils.StringToUintPtr(stringId)
	if err != nil {
		return err
	}

	customer, err := database.GetCustomerByID(*id)
	if err != nil {
		return err
	}

	filter := database.NewInvoiceFilter()
	filter.CustomerID = &customer.ID

	invoiceList, err := database.GetAllInvoice(filter)
	if err != nil {
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	utils.SetPage(echoCtx, enum.CustomerList)

	utils.SetTitle(echoCtx, fmt.Sprintf("%s - customer detail", customer.Name))

	customerDetailParams := CustomerDetailParams{
		Item:        *customer,
		InvoiceList: invoiceList,
		Filter:      filter,
	}

	return utils.Render(CustomerDetail(customerDetailParams), echoCtx)
}

func filterHandler(echoCtx echo.Context) error {

	filter := getCustomerFilterFromContext(echoCtx)

	items, err := database.GetAllCustomerWithTotals(filter)
	if err != nil {
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	return utils.Render(AllCustomerRowsShow(*items), echoCtx)
}

func searchHandler(echoCtx echo.Context) error {
	name := echoCtx.QueryParam("name")

	customerList, err := database.SearchCustomerByName(name)
	if err != nil {
		return err
	}

	return utils.Render(CustomerSearch(*customerList), echoCtx)
}
