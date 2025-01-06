package responses

import "customer_management_service/models"

type InviteHashDTO struct {
	Token *models.UserTokens
}

type InviteHashResponseDTO struct {
	StatusCode int
	Value      *InviteHashDTO
	StatusDesc string
}
