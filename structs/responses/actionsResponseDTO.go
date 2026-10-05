package responses

import "customer_management_service/models"

type ActionResponseDTO struct {
	StatusCode int
	Action     *models.Actions
	StatusDesc string
}

type ActionsResponseDTO struct {
	StatusCode int
	Actions    *[]models.Actions
	StatusDesc string
}

type ActionsAllResponseDTO struct {
	StatusCode int
	Actions    *[]interface{}
	StatusDesc string
}
