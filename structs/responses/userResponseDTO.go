package responses

import (
	"time"
)

type UserResp struct {
	UserId        string
	ImagePath     string
	UserType      string
	FullName      string
	Username      string
	Password      string
	Email         string
	PhoneNumber   string
	Gender        string
	Dob           time.Time
	Address       string
	IdType        string
	IdNumber      string
	MaritalStatus string
	Active        int
	Role          *Roles
	IsVerified    bool
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     string
	ModifiedBy    string
	Branch        *Branches
}

type UserResponseDTO struct {
	StatusCode int
	Result     *UserResp
	StatusDesc string
}

type UsersResponseDTO struct {
	StatusCode int
	Result     *[]UserResp
	StatusDesc string
}

type UsersAllCustomersDTO struct {
	StatusCode int
	Result     *[]UserResp
	StatusDesc string
}

type UsersBranchResponseDTO struct {
	StatusCode int
	Result     *[]UserResp
	StatusDesc string
}
