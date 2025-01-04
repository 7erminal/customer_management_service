package responses

import "customer_management_service/models"

type PermissionResponseDTO struct {
	StatusCode int
	Permission *models.Permissions
	StatusDesc string
}

type PermissionsResponseDTO struct {
	StatusCode  int
	Permissions *[]models.Permissions
	StatusDesc  string
}

type PermissionsAllResponseDTO struct {
	StatusCode  int
	Permissions *[]interface{}
	StatusDesc  string
}
