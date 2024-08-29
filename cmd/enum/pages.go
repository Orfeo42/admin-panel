package enum

type Page string

const (
	Home         Page = "home"
	InvoiceList  Page = "InvoiceList"
	InvoiceAdd   Page = "InvoiceAdd"
	CustomerList Page = "CustomerList"
	CustomerAdd  Page = "CustomerAdd"
	OrderList    Page = "OrderList"
	OrderAdd     Page = "OrderAdd"
)
