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
