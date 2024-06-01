package responses

import "customer_management_service/models"

type NewsLetterCustomerDTO struct {
	StatusCode int
	Customer   *models.Newsletter_customers
	StatusDesc string
}

type NewsLetterCustomersDTO struct {
	StatusCode int
	Customers  *[]models.Newsletter_customers
	StatusDesc string
}

type NewsLetterAllCustomersDTO struct {
	StatusCode int
	Customers  *[]interface{}
	StatusDesc string
}
