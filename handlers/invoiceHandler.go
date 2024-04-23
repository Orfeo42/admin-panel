package handlers

import (
	"math/rand"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/invoice"
	"github.com/labstack/echo/v4"
)

func InvoiceListShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.InvoiceList)
	echoCtx = utils.SetTitle(echoCtx, "Invoices")

	items := []data.InvoiceModel{}

	for i := 0; i < 100; i++ {
		items = append(items, genRandomInvoice())
	}

	return utils.Render(invoice.InvoiceListView(items), echoCtx)
}

func InvoiceShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.InvoiceAdd)
	echoCtx = utils.SetTitle(echoCtx, "Invoice")

	return utils.Render(invoice.InvoiceView(data.InvoiceModel{}), echoCtx)
}

func genRandomInvoice() data.InvoiceModel {
	return data.InvoiceModel{
		Customer: utils.RandomString(25),
		Amount:   rand.Float64(),
		Date:     utils.RandomDate(),
		IsPaid:   rand.Intn(2) == 1,
	}
}
