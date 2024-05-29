package validation

import (
	"errors"

	"admin-panel/mvc/repositories"

	"github.com/labstack/gommon/log"
)

func FindCustomerFromName(customers *[]repositories.Customer, name string) (repositories.Customer, error) {
	for _, customer := range *customers {
		if customer.Name == name {
			return customer, nil
		}
	}
	log.Infof("No customer with name '%s' found", name)
	return repositories.Customer{}, errors.New("No customer with name " + name + " found")
}
