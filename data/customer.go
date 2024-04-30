package data

import (
	"log"

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
	dbInstance, err := db.GetInstance()
	if err != nil {
		return result, err
	}
	dbInstance.Transaction(func(tx *gorm.DB) error {
		for _, customer := range customerList {
			if err := tx.Create(&customer).Error; err != nil {
				log.Fatalln("Error in creating customer!", customer, err)
				return err
			}
			result = append(result, customer)
		}
		return nil
	})

	return result, nil
}
