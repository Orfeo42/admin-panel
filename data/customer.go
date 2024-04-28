package data

import (
	"github.com/Orfeo42/admin-panel/db"
	"gorm.io/gorm"
)

type CustomerModel struct {
	gorm.Model
	Name    string
	Surname string
	Address string
	Phone   string
	Email   string
}

func GetAllCustomer() ([]CustomerModel, error) {

	dbInstance, err := db.GetInstance()
	if err != nil {
		return []CustomerModel{}, err
	}

	var items []CustomerModel
	result := dbInstance.Find(&items)

	return items, result.Error
}

func CreateCustomer(customer CustomerModel) (CustomerModel, error) {

	dbInstance, err := db.GetInstance()
	if err != nil {
		return CustomerModel{}, err
	}

	result := dbInstance.Create(&customer)

	return customer, result.Error
}

func CreateCustomerList(customerList []CustomerModel) ([]CustomerModel, error) {
	result := []CustomerModel{}
	for _, customer := range customerList {
		ret, err := CreateCustomer(customer)
		if err != nil {
			return result, err
		}
		result = append(result, ret)
	}
	return result, nil
}
