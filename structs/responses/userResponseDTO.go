package responses

import "customer_management_service/models"

type UserResponseDTO struct {
	StatusCode int
	User       *models.Users
	StatusDesc string
}
