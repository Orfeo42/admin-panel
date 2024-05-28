package controllers

import (
	"time"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/mvc/services"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/pages/viewinvoice"
	"github.com/labstack/echo/v4"
)

const pageName = "Fatture"

func InvoiceController(application *echo.Echo) {
	invoiceGroup := application.Group("/invoice")

	invoiceGroup.GET("/list", func(echoCtx echo.Context) error {
		filter := getInvoiceFilterFromContext(echoCtx)

		invoices, err := services.GetInvoiceFromFilter(filter)
		if err != nil {
			//TODO GESTIRE L'ERRORE HTML
			return err
		}

		utils.SetPage(echoCtx, enum.InvoiceList)

		utils.SetTitle(echoCtx, pageName)

		utils.SetPageNumber(echoCtx, filter.Page)

		invoiceListParams := viewinvoice.InvoiceListParams{
			Items:  *invoices,
			Filter: filter,
		}

		return utils.Render(viewinvoice.InvoiceList(invoiceListParams), echoCtx)
	})

	invoiceGroup.GET("/filter", func(echoCtx echo.Context) error {

		filter := getInvoiceFilterFromContext(echoCtx)

		invoices, err := services.GetInvoiceFromFilter(filter)
		if err != nil {
			//TODO GESTIRE L'ERRORE HTML
			return err
		}

		utils.SetPageNumber(echoCtx, filter.Page)

		return utils.Render(viewinvoice.InvoiceRows(*invoices), echoCtx)
	})

	invoiceGroup.GET("/add", func(echoCtx echo.Context) error {
		//TODO ESTRARRE DATI
		invoice, err := services.CreateNewInvoice(repositories.Invoice{})
		if err != nil {
			//TODO GESTIRE L'ERRORE HTML
			return err
		}

		utils.SetPage(echoCtx, enum.InvoiceAdd)

		utils.SetTitle(echoCtx, pageName)

		return utils.Render(viewinvoice.InvoiceEdit(*invoice), echoCtx)
	})

	invoiceGroup.GET("/:id/edit", func(echoCtx echo.Context) error {

		id := echoCtx.Param("id")

		inv, err := repositories.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(viewinvoice.InvoiceRowEdit(viewinvoice.InvoiceRowEditParams{
			Item:   *inv,
			Errors: map[string]string{},
		}), echoCtx)
	})

	invoiceGroup.GET("/:id", func(echoCtx echo.Context) error {

		id := echoCtx.Param("id")

		inv, err := repositories.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(viewinvoice.InvoiceRow(*inv), echoCtx)
	})

	invoiceGroup.PUT("/:id", func(echoCtx echo.Context) error {

		id, err := utils.StringToUintPtr(echoCtx.Param("id"))
		if err != nil {
			return err
		}

		invoiceIn, errors := validateCreateUpdateRequest(echoCtx)
		if len(errors) > 0 {
			return utils.Render(viewinvoice.InvoiceRowEdit(viewinvoice.InvoiceRowEditParams{
				Item:   invoiceIn,
				Errors: errors,
			}), echoCtx)
		}

		invoice, err := services.UpdateInvoice(*id, invoiceIn)
		/*if err != nil {
			return err
		}*/

		return utils.Render(viewinvoice.InvoiceRow(*invoice), echoCtx)
	})

	invoiceGroup.PUT("/:id/pay", func(echoCtx echo.Context) error {
		id := echoCtx.Param("id")

		inv, err := repositories.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		inv.PaidAmount = inv.Amount

		paymentDate := time.Now()
		inv.PaymentDate = &paymentDate

		updateInvoice, err := repositories.UpdateInvoice(*inv)
		if err != nil {
			return err
		}

		return utils.Render(viewinvoice.InvoiceRow(*updateInvoice), echoCtx)
	})

	invoiceGroup.POST("", func(echoCtx echo.Context) error {

		invoiceIn, errors := validateCreateUpdateRequest(echoCtx)
		if len(errors) > 0 {
			//TODO GESTISCI ERRORE HTML

		}

		invoice, err := services.CreateNewInvoice(invoiceIn)
		if err != nil {
			return err
		}

		return utils.Render(viewinvoice.InvoiceForm(*invoice), echoCtx)
	})
}

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

func getInvoiceFilterFromContext(echoCtx echo.Context) repositories.InvoiceFilter {

	filter := repositories.NewInvoiceFilter()

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

func validateCreateUpdateRequest(echoCtx echo.Context) (repositories.Invoice, map[string]string) {
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

	return repositories.Invoice{
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
