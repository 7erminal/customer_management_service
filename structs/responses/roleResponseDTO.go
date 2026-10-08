package responses

import (
	"time"
)

type Roles struct {
	RoleId          int64     `orm:"auto"`
	Role            string    `orm:"size(100)"`
	Description     string    `orm:"size(500)"`
	DateCreated     time.Time `orm:"type(datetime);auto_now_add"`
	DateModified    time.Time `orm:"type(datetime);auto_now"`
	CreatedBy       string
	ModifiedBy      string
	Active          int
	RolePermissions []*Role_permissions `orm:"reverse(many)"`
}

type RoleResponseDTO struct {
	StatusCode int
	Role       *Roles
	StatusDesc string
}

type RolesResponseDTO struct {
	StatusCode int
	Roles      *[]Roles
	StatusDesc string
}

type RolesAllResponseDTO struct {
	StatusCode int
	Roles      *[]Roles
	StatusDesc string
}
