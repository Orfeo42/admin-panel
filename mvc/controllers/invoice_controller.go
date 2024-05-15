package controllers

import (
	"time"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func InvoiceController(application *echo.Echo) {
	invoiceGroup := application.Group("/invoice")

	invoiceGroup.GET("/list", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, enum.InvoiceList)
		echoCtx = utils.SetTitle(echoCtx, "Invoices")

		filter := repositories.NewBaseFilter()

		utils.SetPageNumber(echoCtx, 1)
		items, err := repositories.GetAllInvoice(filter)
		if err != nil {
			return err
		}
		return utils.Render(pages.InvoiceListView(items, filter), echoCtx)
	})

	invoiceGroup.GET("/filter", func(echoCtx echo.Context) error {
		customerID, err := utils.StringToUint(echoCtx.FormValue("customer"))
		if err != nil {
			customerID = nil
		}
		filter := repositories.NewBaseFilter()

		page := utils.StringToInt(echoCtx.QueryParam("page"))
		if page != nil {
			filter.Page = *page
			utils.SetPageNumber(echoCtx, *page)
		}

		filter.CustomerID = customerID
		filter.Number = utils.StringToString(echoCtx.FormValue("number"))
		filter.DateFrom = utils.StringToTime(echoCtx.FormValue("dateFrom"))
		filter.DateTo = utils.StringToTime(echoCtx.FormValue("dateTo"))
		filter.PaymentDateFrom = utils.StringToTime(echoCtx.FormValue("paymentDateFrom"))
		filter.PaymentDateTo = utils.StringToTime(echoCtx.FormValue("paymentDateTo"))
		filter.AmountFrom = utils.StringToAmount(echoCtx.FormValue("amountFrom"))
		filter.AmountTo = utils.StringToAmount(echoCtx.FormValue("amountTo"))
		filter.PaidAmountFrom = utils.StringToAmount(echoCtx.FormValue("paidAmountFrom"))
		filter.IsPaid = isPaidToBool(echoCtx.FormValue("isPaid"))

		items, err := repositories.GetAllInvoice(filter)
		if err != nil {
			return err
		}
		return utils.Render(pages.AllInvoiceRowsShow(items), echoCtx)
	})

	invoiceGroup.GET("/add", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, enum.InvoiceAdd)
		echoCtx = utils.SetTitle(echoCtx, "Invoice")

		return utils.Render(pages.InvoiceView(repositories.Invoice{}), echoCtx)
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
		id := echoCtx.Param("id")

		inv, err := repositories.GetInvoiceByIDString(id)
		if err != nil {
			return err
		}

		customerID, err := utils.StringToUint(echoCtx.FormValue("customer"))
		if err != nil {
			customerID = nil
			log.Info("customer convertion error")
		}
		if customerID != nil {
			customer, err := repositories.GetCustomerByID(*customerID)
			if err != nil {
				log.Errorf("CustomerID not valid: %+v", err)
				return err
			}
			inv.CustomerID = *customerID
			inv.Customer = *customer
		}

		number := echoCtx.FormValue("number")
		date := utils.StringToTime(echoCtx.FormValue("date"))
		paymentDate := utils.StringToTime(echoCtx.FormValue("paymentDate"))
		amount := utils.StringToAmount(echoCtx.FormValue("amount"))
		paidAmount := utils.StringToAmount(echoCtx.FormValue("paidAmount"))

		inv.Number = number
		inv.Date = date
		inv.PaymentDate = paymentDate
		zero := 0

		inv.CustomerID = *customerID

		if amount == nil {
			amount = &zero
		}
		inv.Amount = *amount

		if paidAmount == nil {
			paidAmount = &zero
		}
		inv.PaidAmount = *paidAmount

		updateInvoice, err := repositories.UpdateInvoice(*inv)
		if err != nil {
			return err
		}

		return utils.Render(pages.InvoiceRowShow(updateInvoice), echoCtx)
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

		return utils.Render(pages.InvoiceRowShow(updateInvoice), echoCtx)
	})

	invoiceGroup.POST("", func(echoCtx echo.Context) error {

		input := repositories.Invoice{
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
		result, err := repositories.CreateInvoice(input)
		if err != nil {
			return err
		}

		return utils.Render(pages.InvoiceForm(result), echoCtx)
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
