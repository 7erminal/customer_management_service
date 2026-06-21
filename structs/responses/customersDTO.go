package responses

import "customer_management_service/models"

type CustomersDTO struct {
	StatusCode int
	Customers  *[]models.Customers
	StatusDesc string
}
