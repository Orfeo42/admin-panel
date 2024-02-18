package handler

import (
	"time"

	"github.com/Orfeo42/admin-panel/model"
	"github.com/Orfeo42/admin-panel/utils"
	invoice "github.com/Orfeo42/admin-panel/view/page/invoice/list"
	"github.com/labstack/echo/v4"
)

func InvoiceListShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, "Invoices")
	echoCtx = utils.SetTitle(echoCtx, "Invoices")
	invoices := []model.InvoiceModel{
		{Customer: "Customer1", Amount: 10.5, Date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC), IsPaid: false},
		{Customer: "Customer2", Amount: 11.5, Date: time.Date(2022, 4, 2, 0, 0, 0, 0, time.UTC), IsPaid: true},
		{Customer: "Customer3", Amount: 12.5, Date: time.Date(2019, 5, 5, 0, 0, 0, 0, time.UTC), IsPaid: false},
	}
	return utils.Render(invoice.InvoiceView(invoices), echoCtx)
}
