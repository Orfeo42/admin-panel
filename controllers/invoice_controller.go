package controllers

import (
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/invoice"
	"github.com/labstack/echo/v4"
)

func InvoiceController(application *echo.Echo) {
	invoiceGroup := application.Group("/invoice")

	invoiceGroup.GET("/list", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, pages.InvoiceList)
		echoCtx = utils.SetTitle(echoCtx, "Invoices")

		filter := data.NewBaseFilter()

		items, err := data.GetAllInvoice(filter)
		if err != nil {
			return err
		}
		return utils.Render(invoice.InvoiceListView(items, filter), echoCtx)
	})

	invoiceGroup.GET("/filter", func(echoCtx echo.Context) error {
		filter := getFilterFromRequest(echoCtx)
		items, err := data.GetAllInvoice(filter)
		if err != nil {
			return err
		}
		return utils.Render(invoice.AllInvoiceRowsShow(items), echoCtx)
	})

	invoiceGroup.GET("/add", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, pages.InvoiceAdd)
		echoCtx = utils.SetTitle(echoCtx, "Invoice")

		return utils.Render(invoice.InvoiceView(data.Invoice{}), echoCtx)
	})

	invoiceGroup.GET("/:id/edit", func(echoCtx echo.Context) error {
		id := echoCtx.Param("id")

		inv, err := data.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(invoice.InvoiceRowEdit(*inv), echoCtx)
	})

	invoiceGroup.GET("/:id", func(echoCtx echo.Context) error {
		id := echoCtx.Param("id")

		inv, err := data.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(invoice.InvoiceRowShow(*inv), echoCtx)
	})

	invoiceGroup.PUT("/:id", func(echoCtx echo.Context) error {
		id := echoCtx.Param("id")

		inv, err := data.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		return utils.Render(invoice.InvoiceRowEdit(*inv), echoCtx)
	})

	invoiceGroup.POST("", func(echoCtx echo.Context) error {

		input := data.Invoice{
			CustomerID:          0,
			Year:                2022,
			Number:              "",
			PaymentMethod:       nil,
			Amount:              0,
			PaidAmount:          0,
			Date:                nil,
			PaymentDate:         nil,
			ExpectedPaymentDate: nil,
			Rows:                nil,
			Note:                nil,
		}
		result, err := data.CreateInvoice(input)
		if err != nil {
			return err
		}

		return utils.Render(invoice.InvoiceForm(result), echoCtx)
	})
}

func getFilterFromRequest(echoCtx echo.Context) data.InvoiceFilter {
	result := data.InvoiceFilter{
		CustomerID:      utils.StringToUint(echoCtx.FormValue("customer")),
		Number:          utils.StringToString(echoCtx.FormValue("number")),
		DateFrom:        utils.StringToTime(echoCtx.FormValue("dateFrom")),
		DateTo:          utils.StringToTime(echoCtx.FormValue("dateTo")),
		PaymentDateFrom: utils.StringToTime(echoCtx.FormValue("paymentDateFrom")),
		PaymentDateTo:   utils.StringToTime(echoCtx.FormValue("paymentDateTo")),
		AmountFrom:      utils.StringToInt(echoCtx.FormValue("amountFrom")),
		AmountTo:        utils.StringToInt(echoCtx.FormValue("amountTo")),
		PaidAmountFrom:  utils.StringToInt(echoCtx.FormValue("paidAmountFrom")),
		PaidAmountTo:    utils.StringToInt(echoCtx.FormValue("paidAmountTo")),
		IsPaid:          isPaidToBool(echoCtx.FormValue("isPaid")),
	}

	return result
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