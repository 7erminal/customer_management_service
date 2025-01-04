package requests

import "customer_management_service/models"

type RolePermissionRequest struct {
	Role       *models.Roles
	Permission *models.Permissions
}
