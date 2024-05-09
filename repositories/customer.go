package repositories

import (
	"github.com/labstack/gommon/log"

	"github.com/Orfeo42/admin-panel/db"
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

func GetAllCustomer() ([]Customer, error) {
	dbInstance, err := db.GetInstance()
	if err != nil {
		return []Customer{}, err
	}

	var items []Customer
	result := dbInstance.Find(&items)

	return items, result.Error
}

func GetAllCustomerWithTotals() (*[]CustomerWithTotals, error) {

	dbInstance, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	var results []CustomerWithTotals

	response := dbInstance.Table("invoices").
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
		Order("max(customers.name) asc").
		Scan(&results)
	if response.Error != nil {
		return nil, response.Error
	}
	return &results, response.Error
}

func CreateCustomer(customer Customer) (Customer, error) {

	dbInstance, err := db.GetInstance()
	if err != nil {
		return Customer{}, err
	}

	result := dbInstance.Create(&customer)

	return customer, result.Error
}

func CreateCustomerList(customerList *[]Customer) (*[]Customer, error) {
	var result []Customer
	dbInstance, err := db.GetInstance()
	if err != nil {
		return nil, err
	}
	err = dbInstance.Transaction(func(tx *gorm.DB) error {
		for _, customer := range *customerList {
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

	return &result, nil
}

func GetCustomerByID(id uint) (*Customer, error) {
	dbInstance, err := db.GetInstance()
	if err != nil {
		return nil, err
	}
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

	dbInstance, err := db.GetInstance()
	if err != nil {
		return nil, err
	}
	var customerList []Customer
	tx := dbInstance.
		Where("lower(customers.name) LIKE '%' || lower(?) || '%'", name).
		Limit(10).
		Find(&customerList)

	if tx.Error != nil {
		log.Errorf("Error in searching customer by name: %+v", tx.Error)
		return nil, tx.Error
	}
	return &customerList, nil
}
