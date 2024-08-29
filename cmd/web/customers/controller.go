package customers

import (
	"fmt"

	"admin-panel/cmd/enum"
	"admin-panel/cmd/web/invoices"
	"admin-panel/internal/database"
	"admin-panel/utils"

	"github.com/labstack/echo/v4"
)

func RegisterCustomerRoutes(application *echo.Echo) {

	customerGroup := application.Group("/customer")
	controller := getControllerInstance()

	customerGroup.GET("", controller.readAllPage)
	customerGroup.GET("/:id", controller.readPage)
	customerGroup.GET("/filter", controller.filter)
	customerGroup.GET("/search", controller.searchByName)

	customerGroup.GET("/add", controller.createPage)
	customerGroup.POST("", controller.create)

	customerGroup.GET("/:id/edit", controller.updatePage)
	customerGroup.PUT("/:id", controller.update)
}

const pageName = "Fatture"

var controllerInstance controller

type controller interface {
	readAllPage(echoCtx echo.Context) error
	readPage(echoCtx echo.Context) error
	filter(echoCtx echo.Context) error
	searchByName(echoCtx echo.Context) error

	createPage(echoCtx echo.Context) error
	create(echoCtx echo.Context) error

	updatePage(echoCtx echo.Context) error
	update(echoCtx echo.Context) error
}

type customerController struct {
	custRep database.CustomerRepository
	invRep  database.InvoiceRepository
}

func getControllerInstance() controller {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &customerController{
		custRep: database.CustomerRepositoryInstance(),
		invRep:  database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

/*--READ HANDLER--*/

func (c *customerController) readAllPage(echoCtx echo.Context) error {

	filter := getCustomerFilterFromContext(echoCtx)

	items, err := c.custRep.ReadAllFilteredWithTotals(filter)
	if err != nil {
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	utils.SetPage(echoCtx, enum.CustomerList)

	utils.SetTitle(echoCtx, pageName)

	customerListParams := CustomerListParams{
		Items:  items,
		Filter: filter,
	}

	return utils.Render(CustomerListView(customerListParams), echoCtx)
}

func (c *customerController) readPage(echoCtx echo.Context) error {

	stringId := echoCtx.Param("id")

	id, err := utils.StringToUintPtr(stringId)
	if err != nil {
		return err
	}

	customer, err := c.custRep.Read(*id)
	if err != nil {
		return err
	}

	filter := database.NewInvoiceFilter()
	filter.CustomerID = &customer.ID

	invoiceList, err := c.invRep.ReadAllFiltered(filter)
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

func (c *customerController) filter(echoCtx echo.Context) error {
	filter := getCustomerFilterFromContext(echoCtx)

	items, err := c.custRep.ReadAllFilteredWithTotals(filter)
	if err != nil {
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	return utils.Render(AllCustomerRowsShow(items), echoCtx)
}

func (c *customerController) searchByName(echoCtx echo.Context) error {
	name := echoCtx.QueryParam("name")

	customerList, err := c.custRep.ReadAllByName(name)
	if err != nil {
		return err
	}

	return utils.Render(CustomerSearch(customerList), echoCtx)
}

/*--CREATE HANDLER--*/

func (c *customerController) createPage(echoCtx echo.Context) error {
	return nil
}

func (c *customerController) create(echoCtx echo.Context) error {
	return nil
}

/*--UPDATE HANDLER--*/

func (c *customerController) updatePage(echoCtx echo.Context) error {
	return nil
}

func (c *customerController) update(echoCtx echo.Context) error {
	return nil
}

/**/

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
	filter.IsPaid = invoices.IsPaidToBool(echoCtx.FormValue("isPaid"))
	return filter

}
