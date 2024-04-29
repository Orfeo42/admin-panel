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

	items := []data.Invoice{}

	for i := 0; i < 100; i++ {
		items = append(items, genRandomInvoice())
	}

	return utils.Render(invoice.InvoiceListView(items), echoCtx)
}

func InvoiceShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.InvoiceAdd)
	echoCtx = utils.SetTitle(echoCtx, "Invoice")

	return utils.Render(invoice.InvoiceView(data.Invoice{}), echoCtx)
}

func genRandomInvoice() data.Invoice {
	return data.Invoice{
		Customer: data.Customer{
			Name: utils.RandomString(25),
		},
		Amount:      rand.Int(),
		Date:        utils.RandomDate(),
		PaymentDate: utils.RandomDate(),
		PaidAmount:  rand.Int(),
	}
}
