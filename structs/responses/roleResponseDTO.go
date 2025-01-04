package responses

import "customer_management_service/models"

type RoleResponseDTO struct {
	StatusCode int
	Role       *models.Roles
	StatusDesc string
}

type RolesResponseDTO struct {
	StatusCode int
	Roles      *[]models.Roles
	StatusDesc string
}

type RolesAllResponseDTO struct {
	StatusCode int
	Roles      *[]interface{}
	StatusDesc string
}
