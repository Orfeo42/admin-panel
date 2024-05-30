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
	controller := CustomerControllerInstance()

	customerGroup.GET("", controller.ReadAllPageHandler)
	customerGroup.GET("/:id", controller.ReadPageHandler)
	customerGroup.GET("/filter", controller.FilterHandler)
	customerGroup.GET("/search", controller.SearchByNameHandler)

	customerGroup.GET("/add", controller.CreatePageHandler)
	customerGroup.POST("", controller.CreateHandler)

	customerGroup.GET("/:id/edit", controller.UpdatePageHandler)
	customerGroup.PUT("/:id", controller.UpdateHandler)
}

const pageName = "Fatture"

var controllerInstance CustomerController

type CustomerController interface {
	ReadAllPageHandler(echoCtx echo.Context) error
	ReadPageHandler(echoCtx echo.Context) error
	FilterHandler(echoCtx echo.Context) error
	SearchByNameHandler(echoCtx echo.Context) error

	CreatePageHandler(echoCtx echo.Context) error
	CreateHandler(echoCtx echo.Context) error

	UpdatePageHandler(echoCtx echo.Context) error
	UpdateHandler(echoCtx echo.Context) error
}

type customerController struct {
	custRep database.CustomerRepository
	invRep  database.InvoiceRepository
}

func CustomerControllerInstance() CustomerController {
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

func (c *customerController) ReadAllPageHandler(echoCtx echo.Context) error {

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

func (c *customerController) ReadPageHandler(echoCtx echo.Context) error {

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

func (c *customerController) FilterHandler(echoCtx echo.Context) error {
	filter := getCustomerFilterFromContext(echoCtx)

	items, err := c.custRep.ReadAllFilteredWithTotals(filter)
	if err != nil {
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	return utils.Render(AllCustomerRowsShow(items), echoCtx)
}

func (c *customerController) SearchByNameHandler(echoCtx echo.Context) error {
	name := echoCtx.QueryParam("name")

	customerList, err := c.custRep.ReadAllByName(name)
	if err != nil {
		return err
	}

	return utils.Render(CustomerSearch(customerList), echoCtx)
}

/*--CREATE HANDLER--*/

func (c *customerController) CreatePageHandler(echoCtx echo.Context) error {
	return nil
}

func (c *customerController) CreateHandler(echoCtx echo.Context) error {
	return nil
}

/*--UPDATE HANDLER--*/

func (c *customerController) UpdatePageHandler(echoCtx echo.Context) error {
	return nil
}

func (c *customerController) UpdateHandler(echoCtx echo.Context) error {
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
