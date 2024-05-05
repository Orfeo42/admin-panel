package data

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
	var result []Customer
	dbInstance, err := db.GetInstance()
	if err != nil {
		return result, err
	}
	err = dbInstance.Transaction(func(tx *gorm.DB) error {
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
