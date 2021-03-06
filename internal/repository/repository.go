package repository

import (
	"github.com/ekateryna-tln/wallester-task/internal/models"
	"github.com/gofrs/uuid"
)

type DatabaseRepo interface {
	GetAllCustomers() ([]models.Customer, error)
	SearchCustomers(string) ([]models.Customer, error)
	InsertCustomer(customer models.Customer) (string, error)
	GetCustomerByID(uuid uuid.UUID) (models.Customer, error)
	UpdateCustomer(customer models.Customer) error
}
