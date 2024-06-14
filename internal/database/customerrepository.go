package database

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var custRepoInstance CustomerRepository

type CustomerRepository interface {
	Create(customer Customer) (*Customer, error)
	CreateListInTransaction(customerList []Customer) ([]Customer, error)
	CreateList(customerList []Customer) ([]Customer, error)
	Read(id uint) (*Customer, error)
	ReadAll() ([]Customer, error)
	ReadAllFilteredWithTotals(filter CustomerFilter) ([]CustomerWithTotals, error)
	ReadAllByName(name string) ([]Customer, error)
	Update(customer Customer) error
	Delete(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func CustomerRepositoryInstance() CustomerRepository {
	if custRepoInstance != nil {
		return custRepoInstance
	}
	custRepoInstance = &customerRepository{
		db: DBInstance(),
	}
	return custRepoInstance
}

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

func (r *customerRepository) Create(customer Customer) (*Customer, error) {
	result := r.db.Create(&customer)
	return &customer, result.Error
}

func (r *customerRepository) CreateListInTransaction(customerList []Customer) ([]Customer, error) {
	var result []Customer
	err := r.db.Transaction(func(tx *gorm.DB) error {
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

func (r *customerRepository) CreateList(customerList []Customer) ([]Customer, error) {
	var result []Customer
	for _, customer := range customerList {
		if _, err := r.Create(customer); err != nil {
			log.Errorf("Error in creating customer %+v: %+v", customer, err)
			continue
		}
		result = append(result, customer)
	}
	return result, nil
}

func (r *customerRepository) Read(id uint) (*Customer, error) {
	var customer Customer
	tx := r.db.First(&customer, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &customer, nil
}

func (r *customerRepository) ReadAll() ([]Customer, error) {
	var items []Customer
	result := r.db.Find(&items)
	return items, result.Error
}

func (r *customerRepository) ReadAllFilteredWithTotals(filter CustomerFilter) ([]CustomerWithTotals, error) {

	//log.Printf("Filer: %+v", filter)
	var results []CustomerWithTotals

	queryDB := r.db.Table("invoices").
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
	return results, response.Error
}

func (r *customerRepository) ReadAllByName(name string) ([]Customer, error) {
	if name == "" {
		return nil, nil
	}

	var customerList []Customer
	tx := r.db.
		Where("lower(customers.name) LIKE '%' || lower(?) || '%'", name).
		Limit(4).
		Find(&customerList)

	if tx.Error != nil {
		//log.Errorf("Error in searching customer by name: %+v", tx.Error)
		return nil, tx.Error
	}
	return customerList, nil
}

func (r *customerRepository) Update(customer Customer) error {
	result := r.db.Save(&customer)
	return result.Error
}

func (r *customerRepository) Delete(id uint) error {
	var customer Customer
	if err := r.db.First(&customer, 1).Error; err != nil {
		fmt.Println("Record not found!")
		return err
	}

	if err := r.db.Delete(customer, 1).Error; err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}
	return nil
}
