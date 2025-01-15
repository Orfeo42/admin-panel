package database

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var invRepoInstance InvoiceRepository

type InvoiceRepository interface {
	Create(invoice Invoice) (*Invoice, error)
	CreateListInTransaction(invoiceList []Invoice) ([]Invoice, error)
	CreateList(invoiceList []Invoice) ([]Invoice, error)
	Read(id uint) (*Invoice, error)
	ReadAll() ([]Invoice, error)
	ReadAllFiltered(filter InvoiceFilter) ([]Invoice, error)
	Update(invoice Invoice) error
	Delete(id uint) error
	ReadAllByCustomerIDAndPaid(customerID uint, isPaid *bool) ([]Invoice, error)
	SalesByMonth(dateFrom, dateTo time.Time) ([]MoneyByMonthResult, error)
	CollectedByMonth(dateFrom, dateTo time.Time) ([]MoneyByMonthResult, error)
	SalesTotal(dateFrom, dateTo time.Time) (int64, error)
	CollectedTotal(dateFrom, dateTo time.Time) (int64, error)
	ToBeCollectedTotal(dateFrom, dateTo time.Time) (int64, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func InvoiceRepositoryInstance() InvoiceRepository {
	if invRepoInstance != nil {
		return invRepoInstance
	}
	invRepoInstance = &invoiceRepository{
		db: DBInstance(),
	}
	return invRepoInstance
}

type Invoice struct {
	gorm.Model
	CustomerID          uint
	Customer            Customer
	Year                int    //`gorm:"uniqueIndex:year_number"`
	Number              string //`gorm:"uniqueIndex:year_number"`
	PaymentMethod       *string
	Amount              int
	PaidAmount          int
	Date                *time.Time
	PaymentDate         *time.Time
	ExpectedPaymentDate *time.Time
	Rows                *[]InvoiceRow
	Note                *string
}

type InvoiceRow struct {
	gorm.Model
	InvoiceID uint
	Invoice   Invoice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Number    int
	Amount    int
}

type InvoiceFilter struct {
	CustomerID      *uint
	Number          *string
	DateFrom        *time.Time
	DateTo          *time.Time
	PaymentDateFrom *time.Time
	PaymentDateTo   *time.Time
	AmountFrom      *int
	AmountTo        *int
	PaidAmountFrom  *int
	PaidAmountTo    *int
	IsPaid          *bool
	PageSize        int
	Page            int
}

func NewInvoiceFilter() InvoiceFilter {
	return InvoiceFilter{
		Page:     1,
		PageSize: 10,
	}
}

type MoneyByMonthResult struct {
	Year   int
	Month  int
	Amount int64
}

func (r *invoiceRepository) Create(invoice Invoice) (*Invoice, error) {
	result := r.db.Create(&invoice)
	return &invoice, result.Error
}

func (r *invoiceRepository) CreateListInTransaction(invoiceList []Invoice) ([]Invoice, error) {
	var result []Invoice
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, invoice := range invoiceList {
			if err := tx.Create(&invoice).Error; err != nil {
				log.Errorf("Error in creating invoice %+v: %+v", invoice, err)
				return err
			}
			result = append(result, invoice)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *invoiceRepository) CreateList(invoiceList []Invoice) ([]Invoice, error) {
	var result []Invoice
	for _, invoice := range invoiceList {
		if _, err := r.Create(invoice); err != nil {
			log.Errorf("Error in creating invoice %+v: %+v", invoice, err)
			continue
		}
		result = append(result, invoice)
	}
	return result, nil
}

func (r *invoiceRepository) Read(id uint) (*Invoice, error) {
	var invoice Invoice
	tx := r.db.First(&invoice, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &invoice, nil
}

func (r *invoiceRepository) Update(invoice Invoice) error {
	result := r.db.Save(&invoice)
	return result.Error
}

func (r *invoiceRepository) Delete(id uint) error {
	var invoice Invoice
	if err := r.db.First(&invoice, 1).Error; err != nil {
		fmt.Println("Record not found!")
		return err
	}

	if err := r.db.Delete(invoice, 1).Error; err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}
	return nil
}

func (r *invoiceRepository) ReadAll() ([]Invoice, error) {
	filter := NewInvoiceFilter()
	filter.PageSize = -1
	return r.ReadAllFiltered(filter)

}

func (r *invoiceRepository) ReadAllFiltered(filter InvoiceFilter) ([]Invoice, error) {

	var items []Invoice
	queryDB := r.db.Order("date desc, number desc")
	queryDB.Preload("Customer")
	if filter.CustomerID != nil {
		queryDB.Where("customer_id = ?", *filter.CustomerID)
	}
	if filter.Number != nil {
		queryDB.Where("number = ?", *filter.Number)
	}
	if filter.DateFrom != nil {
		queryDB.Where("date >= ?", *filter.DateFrom)
	}
	if filter.DateTo != nil {
		queryDB.Where("date <= ?", *filter.DateTo)
	}
	if filter.PaymentDateFrom != nil {
		queryDB.Where("payment_date >= ?", *filter.PaymentDateFrom)
	}
	if filter.PaymentDateTo != nil {
		queryDB.Where("payment_date <= ?", *filter.PaymentDateTo)
	}
	if filter.AmountFrom != nil {
		queryDB.Where("amount >= ?", *filter.AmountFrom)
	}
	if filter.AmountTo != nil {
		queryDB.Where("amount <= ?", *filter.AmountTo)
	}
	if filter.PaidAmountFrom != nil {
		queryDB.Where("paid_amount >= ?", *filter.PaidAmountFrom)
	}
	if filter.PaidAmountTo != nil {
		queryDB.Where("paid_amount <= ?", *filter.PaidAmountTo)
	}
	if filter.IsPaid != nil {
		if *filter.IsPaid {
			queryDB.Where("amount = paid_amount")
		}
		if !*filter.IsPaid {
			queryDB.Where("amount <> paid_amount")
		}
	}
	if filter.PageSize > 1 {
		queryDB.Limit(filter.PageSize)
	}
	offset := 0
	if filter.Page > 1 && filter.PageSize > 1 {
		offset = filter.PageSize * filter.Page
	}
	queryDB.Offset(offset)

	result := queryDB.Find(&items)

	return items, result.Error
}

func (r *invoiceRepository) ReadAllByCustomerIDAndPaid(customerID uint, isPaid *bool) ([]Invoice, error) {
	filter := NewInvoiceFilter()
	filter.CustomerID = &customerID
	filter.IsPaid = isPaid
	filter.PageSize = -1
	return r.ReadAllFiltered(filter)
}

func (r *invoiceRepository) SalesByMonth(dateFrom, dateTo time.Time) ([]MoneyByMonthResult, error) {
	var earningsByMonthResult []MoneyByMonthResult

	r.db.Table("invoices").
		Select("date_part('year', date) as Year, date_part('month', date) as Month, sum(amount) as Amount").
		Where("date between ? and ?", dateFrom, dateTo).
		Group("date_part('year', date), date_part('month', date)").
		Order("date_part('year', date), date_part('month', date)").
		Scan(&earningsByMonthResult)
	return earningsByMonthResult, nil
}

func (r *invoiceRepository) CollectedByMonth(dateFrom, dateTo time.Time) ([]MoneyByMonthResult, error) {
	var earningsByMonthResult []MoneyByMonthResult

	r.db.Table("invoices").
		Select("date_part('year', payment_date) as Year, date_part('month', payment_date) as Month, sum(paid_amount) as Amount").
		Where("payment_date between ? and ?", dateFrom, dateTo).
		Group("date_part('year', payment_date), date_part('month', payment_date)").
		Order("date_part('year', payment_date), date_part('month', payment_date)").
		Scan(&earningsByMonthResult)
	return earningsByMonthResult, nil
}

func (r *invoiceRepository) SalesTotal(dateFrom, dateTo time.Time) (int64, error) {
	var totalAmount int64
	r.db.Table("invoices").
		Select("sum(amount) as Amount").
		Where("date between ? and ?", dateFrom, dateTo).
		Scan(&totalAmount)
	return totalAmount, nil
}

func (r *invoiceRepository) CollectedTotal(dateFrom, dateTo time.Time) (int64, error) {
	var totalAmount int64
	r.db.Table("invoices").
		Select("sum(paid_amount) as Amount").
		Where("payment_date between ? and ?", dateFrom, dateTo).
		Scan(&totalAmount)
	return totalAmount, nil
}

func (r *invoiceRepository) ToBeCollectedTotal(dateFrom, dateTo time.Time) (int64, error) {
	totalSales, err := r.SalesTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in query execution: %+v", err)
		return 0, err
	}
	totalCollected, err := r.CollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in query execution: %+v", err)
		return 0, err
	}
	return totalSales - totalCollected, nil
}
