package database

import (
	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Surname string
	Address string
	Phone   string
	Email   string
}

type CustomerWithTotals struct {
	Customer
	TotalAmount int
	TotalToPay  int
}

type CustomerFilter struct {
	Name            *string
	TotalAmountFrom *int
	TotalAmountTo   *int
	TotalToPayFrom  *int
	TotalToPayTo    *int
	IsPaid          *bool
	PageSize        int
	Page            int
}

func NewCustomerFilter() CustomerFilter {
	return CustomerFilter{
		Page:     1,
		PageSize: 10,
	}
}

func GetAllCustomer() ([]Customer, error) {
	dbInstance := DBInstance()
	var items []Customer
	result := dbInstance.Find(&items)

	return items, result.Error
}

func GetAllCustomerWithTotals(filter CustomerFilter) (*[]CustomerWithTotals, error) {

	dbInstance := DBInstance()

	var results []CustomerWithTotals

	queryDB := dbInstance.Table("invoices").
		Select("invoices.customer_id as id," +
			"max(customers.name) as name," +
			"max(customers.surname) as surname," +
			"max(customers.address) as address," +
			"max(customers.phone) as phone," +
			"max(customers.email) as email," +
			"sum(invoices.amount) as total_amount," +
			"(sum(invoices.amount) - sum(invoices.paid_amount)) as total_to_pay").
		Joins("left outer join customers on customers.id = invoices.customer_id").
		Group("invoices.customer_id").
		Order("max(customers.name) asc")

	if filter.Name != nil {
		queryDB.Having("lower(max(customers.name)) LIKE '%' || lower(?) || '%'", *filter.Name)
	}
	if filter.TotalAmountFrom != nil {
		queryDB.Having("sum(invoices.amount) >= ?", *filter.TotalAmountFrom)
	}
	if filter.TotalAmountTo != nil {
		queryDB.Having("sum(invoices.amount) <= ?", *filter.TotalAmountTo)
	}
	if filter.TotalToPayFrom != nil {
		queryDB.Having("(sum(invoices.amount) - sum(invoices.paid_amount)) >= ?", *filter.TotalToPayFrom)
	}
	if filter.TotalToPayTo != nil {
		queryDB.Having("(sum(invoices.amount) - sum(invoices.paid_amount)) <= ?", *filter.TotalToPayTo)
	}
	if filter.IsPaid != nil {
		if *filter.IsPaid {
			queryDB.Having("sum(invoices.amount) = sum(invoices.paid_amount)")
		} else {
			queryDB.Having("sum(invoices.amount) <> sum(invoices.paid_amount)")
		}
	}
	queryDB.Limit(filter.PageSize)
	offset := 0
	if filter.Page > 1 {
		offset = filter.PageSize * filter.Page
	}
	queryDB.Offset(offset)

	response := queryDB.Scan(&results)
	if response.Error != nil {
		return nil, response.Error
	}
	return &results, response.Error
}

func CreateCustomer(customer Customer) (Customer, error) {

	dbInstance := DBInstance()

	result := dbInstance.Create(&customer)

	return customer, result.Error
}

func CreateCustomerList(customerList []Customer) ([]Customer, error) {
	var result []Customer
	dbInstance := DBInstance()
	err := dbInstance.Transaction(func(tx *gorm.DB) error {
		for _, customer := range customerList {
			if err := tx.Create(&customer).Error; err != nil {
				log.Errorf("Error in creating customer %+v: %+v", customer, err)
				return err
			}
			result = append(result, customer)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetCustomerByID(id uint) (*Customer, error) {
	dbInstance := DBInstance()
	var customer Customer
	tx := dbInstance.First(&customer, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &customer, nil
}

func SearchCustomerByName(name string) (*[]Customer, error) {
	if name == "" {
		return &[]Customer{}, nil
	}

	dbInstance := DBInstance()

	var customerList []Customer
	tx := dbInstance.
		Where("lower(customers.name) LIKE '%' || lower(?) || '%'", name).
		Limit(4).
		Find(&customerList)

	if tx.Error != nil {
		log.Errorf("Error in searching customer by name: %+v", tx.Error)
		return nil, tx.Error
	}
	return &customerList, nil
}
