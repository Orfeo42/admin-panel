package invoices

import (
	"net/http"
	"time"

	"admin-panel/cmd/enum"
	"admin-panel/internal/database"
	"admin-panel/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const baseUrl = "/invoice"

func RegisterRoutes(application *echo.Echo) {
	invoiceGroup := application.Group(baseUrl)
	controller := getControllerInstance()

	invoiceGroup.GET("", controller.readAllPage)
	invoiceGroup.GET("/:id/row", controller.readRow)
	invoiceGroup.GET("/:id", controller.readPage)
	invoiceGroup.GET("/filter", controller.filter)

	invoiceGroup.GET("/add", controller.createPage)
	invoiceGroup.POST("", controller.create)

	invoiceGroup.GET("/:id/edit", controller.updatePage)
	invoiceGroup.PUT("/:id", controller.update)
	invoiceGroup.PUT("/:id/pay", controller.payByID) //TODO MEGLIO PATHC
	invoiceGroup.DELETE("/:id", controller.delete)
}

const pageName = "Fatture"

type Controller interface {
	readAllPage(echoCtx echo.Context) error
	readRow(echoCtx echo.Context) error
	readPage(echoCtx echo.Context) error
	filter(echoCtx echo.Context) error

	createPage(echoCtx echo.Context) error
	create(echoCtx echo.Context) error

	updatePage(echoCtx echo.Context) error
	update(echoCtx echo.Context) error
	payByID(echoCtx echo.Context) error

	delete(echoCtx echo.Context) error
}

type controller struct {
	invRep database.InvoiceRepository
}

var controllerInstance *controller

func getControllerInstance() Controller {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &controller{
		invRep: database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

/*--READ HANDLER--*/

func (c *controller) readAllPage(echoCtx echo.Context) error {
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

func (c *controller) readRow(echoCtx echo.Context) error {
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

func (c *controller) readPage(echoCtx echo.Context) error {
	id := echoCtx.Param("id")
	idUint, err := utils.StringToUint(id)
	if err != nil {
		return err
	}
	invoice, err := c.invRep.Read(idUint)
	if err != nil {
		return err
	}
	///TODO FARE PAGINA DETTAGLIO
	return utils.Render(InvoiceTableRow(*invoice), echoCtx)
}

func (c *controller) filter(echoCtx echo.Context) error {

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

func (c *controller) createPage(echoCtx echo.Context) error {

	utils.SetPage(echoCtx, enum.InvoiceAdd)

	utils.SetTitle(echoCtx, pageName)

	return utils.Render(InvoiceRowAdd(InvoiceEditParams{
		Invoice: database.Invoice{},
		Errors:  nil,
	}), echoCtx)
}

func (c *controller) create(echoCtx echo.Context) error {

	invoiceIn, errors := validateCreateUpdateRequest(echoCtx)

	if len(errors) > 0 {
		return utils.Render(InvoiceRowAdd(InvoiceEditParams{
			Invoice: invoiceIn,
			Errors:  errors,
		}), echoCtx)
	}

	_, err := c.invRep.Create(invoiceIn)
	if err != nil {
		return err
	}
	return echoCtx.Redirect(http.StatusMovedPermanently, baseUrl)
}

/*--UPDATE HANDLER--*/

func (c *controller) updatePage(echoCtx echo.Context) error {
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

func (c *controller) update(echoCtx echo.Context) error {

	id, err := utils.StringToUint(echoCtx.Param("id"))
	if err != nil {
		//TODO GESTISCI ERRORE HTML
		return err
	}

	invoice, errors := validateCreateUpdateRequest(echoCtx)
	invoice.ID = id

	log.Errorf("errors: %+v", errors)
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
	invoiceRes, err := c.invRep.Read(id)
	if err != nil {
		return err
	}

	return utils.Render(InvoiceTableRow(*invoiceRes), echoCtx)
}

func (c *controller) payByID(echoCtx echo.Context) error {
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
func (c *controller) delete(echoCtx echo.Context) error {
	id, err := utils.StringToUint(echoCtx.Param("id"))
	if err != nil {
		//TODO GESTISCI ERRORE HTML

		log.Errorf("errors: %+v", err)
		return err
	}
	err = c.invRep.Delete(id)
	if err != nil {
		//TODO GESTISCI ERRORE HTML
		log.Errorf("errors: %+v", err)
		return err
	}
	return nil
}

/*--VARIUS FUNCTIONS--*/

func IsPaidToBool(valueFrom string) *bool {
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
	filter.IsPaid = IsPaidToBool(echoCtx.FormValue("isPaid"))
	return filter

}

func validateCreateUpdateRequest(echoCtx echo.Context) (database.Invoice, map[string]string) {
	errors := map[string]string{}
	invoice := database.Invoice{}
	var err error
	invoice.CustomerID, err = utils.StringToUint(echoCtx.FormValue("customer"))
	if err != nil {
		errors["customer"] = "customer is not valid"
	}

	invoice.Number = echoCtx.FormValue("number")
	if invoice.Number == "" {
		errors["number"] = "number is not valid"
	}

	invoice.Date, err = utils.StringToTimePtr(echoCtx.FormValue("date"))
	if err != nil {
		errors["date"] = "date is not valid"
	}

	invoice.PaymentDate, err = utils.StringToTimePtr(echoCtx.FormValue("paymentDate"))
	if err != nil {
		errors["paymentDate"] = "payment date is not valid"
	}

	invoice.Amount, err = utils.StringToAmount(echoCtx.FormValue("amount"))
	if err != nil {
		errors["amount"] = "amount is not valid"
	}

	invoice.PaidAmount, err = utils.StringToAmountValidEmpty(echoCtx.FormValue("paidAmount"))
	if err != nil {
		errors["paidAmount"] = "paid amount is not valid"
	}

	invoice.PaymentMethod = utils.StringPtrNilIfEmpty(echoCtx.FormValue("paymentMethod"))

	invoice.Note = utils.StringPtrNilIfEmpty(echoCtx.FormValue("note"))

	return invoice, errors
}
