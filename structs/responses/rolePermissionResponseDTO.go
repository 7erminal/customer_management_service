package responses

import "customer_management_service/models"

type RolePermissionResponseDTO struct {
	StatusCode     int
	RolePermission *models.Role_permissions
	StatusDesc     string
}

type RolePermissionsResponseDTO struct {
	StatusCode      int
	RolePermissions *[]models.Role_permissions
	StatusDesc      string
}

type RolePermissionsAllResponseDTO struct {
	StatusCode      int
	RolePermissions *[]interface{}
	StatusDesc      string
}
