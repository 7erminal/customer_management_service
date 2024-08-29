package responses

import "customer_management_service/models"

type UserResponseDTO struct {
	StatusCode int
	User       *models.Users
	StatusDesc string
}

type UsersResponseDTO struct {
	StatusCode int
	Users      *[]models.Users
	StatusDesc string
}

type UsersAllCustomersDTO struct {
	StatusCode int
	Users      *[]interface{}
	StatusDesc string
}
