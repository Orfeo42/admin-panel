package handlers

import (
	"encoding/json"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/invoice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func InvoiceListShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.InvoiceList)
	echoCtx = utils.SetTitle(echoCtx, "Invoices")

	filter := data.NewBaseFilter()

	items, err := data.GetAllInvoice(filter)
	if err != nil {
		return err
	}

	return utils.Render(invoice.InvoiceListView(items, filter), echoCtx)
}

func InvoiceFilter(echoCtx echo.Context) error {
	filter := getFilterFromRequest(echoCtx)
	items, err := data.GetAllInvoice(filter)
	if err != nil {
		return err
	}

	return utils.Render(invoice.AllInvoiceRowsShow(items), echoCtx)
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

func InvoiceShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.InvoiceAdd)
	echoCtx = utils.SetTitle(echoCtx, "Invoice")

	return utils.Render(invoice.InvoiceView(data.Invoice{}), echoCtx)
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	return jsonBody
}
