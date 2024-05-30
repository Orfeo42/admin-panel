package validation

import (
	"admin-panel/internal/database"
	"errors"

	"github.com/labstack/gommon/log"
)

func FindCustomerFromName(customers []database.Customer, name string) (database.Customer, error) {
	for _, customer := range customers {
		if customer.Name == name {
			return customer, nil
		}
	}
	log.Infof("No customer with name '%s' found", name)
	return database.Customer{}, errors.New("No customer with name " + name + " found")
}
