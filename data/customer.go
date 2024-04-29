package data

import (
	"github.com/Orfeo42/admin-panel/db"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name    string
	Surname string
	Address string
	Phone   string
	Email   string
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

func CreateCustomer(customer Customer) (Customer, error) {

	dbInstance, err := db.GetInstance()
	if err != nil {
		return Customer{}, err
	}

	result := dbInstance.Create(&customer)

	return customer, result.Error
}

func CreateCustomerList(customerList []Customer) ([]Customer, error) {
	result := []Customer{}
	for _, customer := range customerList {
		ret, err := CreateCustomer(customer)
		if err != nil {
			return result, err
		}
		result = append(result, ret)
	}
	return result, nil
}
