package requests

type RolePermissionRequest struct {
	Role           string
	Action         string
	PermissionCode string
	AddedBy        string
}

type RemoveRolePermissionRequest struct {
	Role           string
	Action         string
	PermissionCode string
	RemovedBy      string
}
