package controllers

import (
	"errors"
	"time"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/mvc/services"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/pages"
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

		return utils.Render(pages.InvoiceListView(*invoices, filter), echoCtx)
	})

	invoiceGroup.GET("/filter", func(echoCtx echo.Context) error {

		filter := getInvoiceFilterFromContext(echoCtx)

		invoices, err := services.GetInvoiceFromFilter(filter)
		if err != nil {
			//TODO GESTIRE L'ERRORE HTML
			return err
		}

		utils.SetPageNumber(echoCtx, filter.Page)

		return utils.Render(pages.AllInvoiceRowsShow(*invoices), echoCtx)
	})

	invoiceGroup.GET("/add", func(echoCtx echo.Context) error {
		//TODO ESTRARRE DATI
		invoice, err := services.CreateNewInvoice(services.InvoiceDTO{})
		if err != nil {
			//TODO GESTIRE L'ERRORE HTML
			return err
		}

		utils.SetPage(echoCtx, enum.InvoiceAdd)

		utils.SetTitle(echoCtx, pageName)

		return utils.Render(pages.InvoiceView(*invoice), echoCtx)
	})

	invoiceGroup.GET("/:id/edit", func(echoCtx echo.Context) error {

		id := echoCtx.Param("id")

		inv, err := repositories.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(pages.InvoiceRowEdit(*inv), echoCtx)
	})

	invoiceGroup.GET("/:id", func(echoCtx echo.Context) error {

		id := echoCtx.Param("id")

		inv, err := repositories.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(pages.InvoiceRowShow(*inv), echoCtx)
	})

	invoiceGroup.PUT("/:id", func(echoCtx echo.Context) error {

		id, err := utils.StringToUintPtr(echoCtx.Param("id"))
		if err != nil {
			//TODO GESTISCI ERRORE HTML
			return errors.New("invoice id is not valid")
		}

		invoiceDTO, err := getInvoiceDTOFromContext(echoCtx)
		if err != nil {
			//TODO GESTISCI ERRORE HTML
			return err
		}

		invoice, err := services.UpdateInvoice(*id, *invoiceDTO)
		if err != nil {
			return err
		}

		return utils.Render(pages.InvoiceRowShow(*invoice), echoCtx)
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

		return utils.Render(pages.InvoiceRowShow(*updateInvoice), echoCtx)
	})

	invoiceGroup.POST("", func(echoCtx echo.Context) error {

		invoiceDTO, err := getInvoiceDTOFromContext(echoCtx)
		if err != nil {
			return err
		}

		invoice, err := services.CreateNewInvoice(*invoiceDTO)
		if err != nil {
			return err
		}

		return utils.Render(pages.InvoiceForm(*invoice), echoCtx)
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

func getInvoiceDTOFromContext(echoCtx echo.Context) (*services.InvoiceDTO, error) {

	customer, err := utils.StringToUintPtr(echoCtx.FormValue("customer"))
	if err != nil {
		return nil, errors.New("customer is not valid")
	}

	number := echoCtx.FormValue("number")
	if number == "" {
		return nil, errors.New("number is not valid")
	}

	date, err := utils.StringToTimePtr(echoCtx.FormValue("date"))
	if err != nil {
		return nil, errors.New("date is not valid")
	}

	paymentDate, err := utils.StringToTimePtr(echoCtx.FormValue("paymentDate"))
	if err != nil {
		return nil, errors.New("payment date is not valid")
	}

	amount, err := utils.StringToAmountPtr(echoCtx.FormValue("amount"))
	if err != nil {
		return nil, errors.New("amount is not valid")
	}

	paidAmount, err := utils.StringToAmountPtr(echoCtx.FormValue("paidAmount"))
	if err != nil {
		return nil, errors.New("paid amount is not valid")
	}

	/*paidAmount, err := utils.StringToAmountPtr(echoCtx.FormValue("paidAmount"))
	if err != nil {
		return nil, errors.New("paid amount is not valid")
	}*/

	return &services.InvoiceDTO{
		CustomerID:          *customer,
		Year:                0,
		Number:              number,
		PaymentMethod:       nil,
		Amount:              *amount,
		PaidAmount:          *paidAmount,
		Date:                date,
		PaymentDate:         paymentDate,
		ExpectedPaymentDate: nil,
		Note:                nil,
	}, nil
}
