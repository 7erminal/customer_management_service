package responses

import "customer_management_service/models"

type UserExtraDetailsResponseDTO struct {
	StatusCode  int
	UserDetails *models.UserExtraDetails
	StatusDesc  string
}
