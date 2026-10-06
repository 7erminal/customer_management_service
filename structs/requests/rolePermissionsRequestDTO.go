package requests

type RolePermissionRequest struct {
	Role           string
	Action         string
	PermissionCode string
	AddedBy        string
}
