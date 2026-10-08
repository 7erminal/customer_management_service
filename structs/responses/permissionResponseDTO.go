package responses

import "time"

type Permissions struct {
	PermissionId          string
	Permission            string
	PermissionCode        string
	PermissionDescription string
	DateCreated           time.Time
	DateModified          time.Time
	CreatedBy             string
	ModifiedBy            string
	Active                int
}

type PermissionResponseDTO struct {
	StatusCode int
	Permission *Permissions
	StatusDesc string
}

type PermissionsResponseDTO struct {
	StatusCode  int
	Permissions *[]Permissions
	StatusDesc  string
}

type PermissionsAllResponseDTO struct {
	StatusCode  int
	Permissions *[]interface{}
	StatusDesc  string
}
