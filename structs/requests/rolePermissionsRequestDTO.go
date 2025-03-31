package requests

type RolePermissionRequest struct {
	Role           int64
	Action         string
	PermissionCode string
}
