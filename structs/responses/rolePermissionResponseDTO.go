package responses

import "time"

type Role_permissions struct {
	RolePermissionId string
	Role             *Roles
	Permission       *Permissions
	Action           *Actions
	DateCreated      time.Time
	DateModified     time.Time
	CreatedBy        string
	ModifiedBy       string
	Active           int
}

type RolePermissionResponseDTO struct {
	StatusCode     int
	RolePermission *Role_permissions
	StatusDesc     string
}

type RolePermissionsResponseDTO struct {
	StatusCode      int
	RolePermissions *[]Role_permissions
	StatusDesc      string
}

type RolePermissionsAllResponseDTO struct {
	StatusCode      int
	RolePermissions *[]interface{}
	StatusDesc      string
}
