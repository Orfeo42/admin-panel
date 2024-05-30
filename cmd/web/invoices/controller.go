package invoices

import (
	"time"

	"admin-panel/cmd/enum"
	"admin-panel/internal/database"
	"admin-panel/utils"

	"github.com/labstack/echo/v4"
)

func RegisterInvoiceRoutes(application *echo.Echo) {
	invoiceGroup := application.Group("/invoice")
	controller := InvoiceControllerInstance()

	invoiceGroup.GET("", controller.ReadAllPageHandler)
	invoiceGroup.GET("/:id", controller.ReadPageHandler)
	invoiceGroup.GET("/filter", controller.FilterHandler)

	invoiceGroup.GET("/add", controller.CreatePageHandler)
	invoiceGroup.POST("", controller.CreateHandler)

	invoiceGroup.GET("/:id/edit", controller.UpdatePageHandler)
	invoiceGroup.PUT("/:id", controller.UpdateHandler)
	invoiceGroup.PUT("/:id/pay", controller.PayByID)
}

const pageName = "Fatture"

var controllerInstance *invoiceController

type InvoiceController interface {
	ReadAllPageHandler(echoCtx echo.Context) error
	ReadPageHandler(echoCtx echo.Context) error
	FilterHandler(echoCtx echo.Context) error

	CreatePageHandler(echoCtx echo.Context) error
	CreateHandler(echoCtx echo.Context) error

	UpdatePageHandler(echoCtx echo.Context) error
	UpdateHandler(echoCtx echo.Context) error
	PayByID(echoCtx echo.Context) error
}

type invoiceController struct {
	invRep database.InvoiceRepository
}

func InvoiceControllerInstance() InvoiceController {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &invoiceController{}
	return controllerInstance
}

/*--READ HANDLER--*/

func (c *invoiceController) ReadAllPageHandler(echoCtx echo.Context) error {
	filter := getInvoiceFilterFromContext(echoCtx)

	invoices, err := c.invRep.ReadAllFiltered(filter)
	if err != nil {
		//TODO GESTIRE L'ERRORE HTML
		return err
	}

	utils.SetPage(echoCtx, enum.InvoiceList)

	utils.SetTitle(echoCtx, pageName)

	utils.SetPageNumber(echoCtx, filter.Page)

	invoiceListParams := InvoiceListParams{
		Items:  invoices,
		Filter: filter,
	}

	return utils.Render(InvoiceList(invoiceListParams), echoCtx)
}

func (c *invoiceController) ReadPageHandler(echoCtx echo.Context) error {
	id := echoCtx.Param("id")
	idUint, err := utils.StringToUint(id)
	if err != nil {
		return err
	}
	invoice, err := c.invRep.Read(idUint)
	if err != nil {
		return err
	}

	return utils.Render(InvoiceTableRow(*invoice), echoCtx)
}

func (c *invoiceController) FilterHandler(echoCtx echo.Context) error {

	filter := getInvoiceFilterFromContext(echoCtx)

	invoices, err := c.invRep.ReadAllFiltered(filter)
	if err != nil {
		//TODO GESTIRE L'ERRORE HTML
		return err
	}

	utils.SetPageNumber(echoCtx, filter.Page)

	return utils.Render(InvoiceTableRows(invoices), echoCtx)
}

/*--CREATE HANDLER--*/

func (c *invoiceController) CreatePageHandler(echoCtx echo.Context) error {
	invoiceIn, errors := validateCreateUpdateRequest(echoCtx)
	if len(errors) > 0 {
		return utils.Render(InvoiceEdit(InvoiceEditParams{
			Invoice: invoiceIn,
			Errors:  errors,
		}), echoCtx)
	}
	invoice, err := c.invRep.Create(invoiceIn)
	if err != nil {
		//TODO GESTIRE L'ERRORE HTML
		return err
	}

	utils.SetPage(echoCtx, enum.InvoiceAdd)

	utils.SetTitle(echoCtx, pageName)

	return utils.Render(InvoiceEdit(InvoiceEditParams{
		Invoice: *invoice,
		Errors:  nil,
	}), echoCtx)
}

func (c *invoiceController) CreateHandler(echoCtx echo.Context) error {

	invoiceIn, errors := validateCreateUpdateRequest(echoCtx)
	if len(errors) > 0 {
		return utils.Render(InvoiceForm(InvoiceEditParams{
			Invoice: invoiceIn,
			Errors:  errors,
		}), echoCtx)
	}

	invoice, err := c.invRep.Create(invoiceIn)
	if err != nil {
		return err
	}

	return utils.Render(InvoiceForm(InvoiceEditParams{
		Invoice: *invoice,
		Errors:  nil,
	}), echoCtx)
}

/*--UPDATE HANDLER--*/

func (c *invoiceController) UpdatePageHandler(echoCtx echo.Context) error {
	id := echoCtx.Param("id")
	idUint, err := utils.StringToUint(id)
	if err != nil {
		return err
	}
	invoice, err := c.invRep.Read(idUint)
	if err != nil {
		return err
	}

	return utils.Render(InvoiceRowEdit(InvoiceEditParams{
		Invoice: *invoice,
		Errors:  map[string]string{},
	}), echoCtx)
}

func (c *invoiceController) UpdateHandler(echoCtx echo.Context) error {

	id, err := utils.StringToUint(echoCtx.Param("id"))
	if err != nil {
		//TODO GESTISCI ERRORE HTML
		return err
	}

	invoice, errors := validateCreateUpdateRequest(echoCtx)
	if len(errors) > 0 {
		return utils.Render(InvoiceRowEdit(InvoiceEditParams{
			Invoice: invoice,
			Errors:  errors,
		}), echoCtx)
	}

	invoice.ID = id

	err = c.invRep.Update(invoice)
	if err != nil {
		return err
	}

	return utils.Render(InvoiceTableRow(invoice), echoCtx)
}

func (c *invoiceController) PayByID(echoCtx echo.Context) error {
	id, err := utils.StringToUint(echoCtx.Param("id"))
	if err != nil {
		//TODO GESTISCI ERRORE HTML
		return err
	}

	invoice, err := c.invRep.Read(id)
	if err != nil {
		//TODO GESTISCI ERRORE HTML
		return err
	}

	invoice.PaidAmount = invoice.Amount

	paymentDate := time.Now()
	invoice.PaymentDate = &paymentDate

	err = c.invRep.Update(*invoice)
	if err != nil {
		return err
	}

	return utils.Render(InvoiceTableRow(*invoice), echoCtx)
}

/*--VARIUS FUNCTIONS--*/

func isPaidToBool(valueFrom string) *bool {
	if valueFrom == "" {
		return nil
	}
	value := false
	if valueFrom == "true" {
		value = true
	}
	return &value
}

func getInvoiceFilterFromContext(echoCtx echo.Context) database.InvoiceFilter {

	filter := database.NewInvoiceFilter()

	customerID, err := utils.StringToUintPtr(echoCtx.FormValue("customer"))
	if err != nil {
		customerID = nil
	}

	page, err := utils.StringToInt(echoCtx.QueryParam("page"))
	if err == nil {
		filter.Page = *page
	}

	filter.CustomerID = customerID
	filter.Number = utils.ParseString(echoCtx.FormValue("number"))
	filter.DateFrom = utils.StringToTimePtrNoErr(echoCtx.FormValue("dateFrom"))
	filter.DateTo = utils.StringToTimePtrNoErr(echoCtx.FormValue("dateTo"))
	filter.PaymentDateFrom = utils.StringToTimePtrNoErr(echoCtx.FormValue("paymentDateFrom"))
	filter.PaymentDateTo = utils.StringToTimePtrNoErr(echoCtx.FormValue("paymentDateTo"))
	filter.AmountFrom = utils.StringToAmountPtrNoErr(echoCtx.FormValue("amountFrom"))
	filter.AmountTo = utils.StringToAmountPtrNoErr(echoCtx.FormValue("amountTo"))
	filter.PaidAmountFrom = utils.StringToAmountPtrNoErr(echoCtx.FormValue("paidAmountFrom"))
	filter.IsPaid = isPaidToBool(echoCtx.FormValue("isPaid"))
	return filter

}

func validateCreateUpdateRequest(echoCtx echo.Context) (database.Invoice, map[string]string) {
	errors := map[string]string{}

	customer, err := utils.StringToUintPtr(echoCtx.FormValue("customer"))
	if err != nil {
		errors["customer"] = "customer is not valid"
	}

	number := echoCtx.FormValue("number")
	if number == "" {
		errors["number"] = "number is not valid"
	}

	date, err := utils.StringToTimePtr(echoCtx.FormValue("date"))
	if err != nil {
		errors["date"] = "date is not valid"
	}

	paymentDate, err := utils.StringToTimePtr(echoCtx.FormValue("paymentDate"))
	if err != nil {
		errors["paymentDate"] = "payment date is not valid"
	}

	amount, err := utils.StringToAmountPtr(echoCtx.FormValue("amount"))
	if err != nil {
		errors["amount"] = "amount is not valid"
	}

	paidAmount, err := utils.StringToAmountPtr(echoCtx.FormValue("paidAmount"))
	if err != nil {
		errors["paidAmount"] = "paid amount is not valid"
	}

	paymentMethod := utils.StringPtrNilIfEmpty(echoCtx.FormValue("paymentMethod"))

	expectedPaymentDate, err := utils.StringToTimePtr(echoCtx.FormValue("expectedPaymentDate"))
	if err != nil {
		errors["expectedPaymentDate"] = "payment date is not valid"
	}

	note := utils.StringPtrNilIfEmpty(echoCtx.FormValue("note"))

	return database.Invoice{
		CustomerID:          *customer,
		Year:                0,
		Number:              number,
		PaymentMethod:       paymentMethod,
		Amount:              *amount,
		PaidAmount:          *paidAmount,
		Date:                date,
		PaymentDate:         paymentDate,
		ExpectedPaymentDate: expectedPaymentDate,
		Note:                note,
	}, errors
}
