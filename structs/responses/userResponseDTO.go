package responses

import (
	"customer_management_service/models"
	"time"
)

type UserResp struct {
	UserId        int64
	ImagePath     string
	UserType      int
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
	Role          *models.Roles
	IsVerified    bool
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
	Branch        *models.Branches
}

type UserResponseDTO struct {
	StatusCode int
	User       *models.Users
	StatusDesc string
}

type UsersResponseDTO struct {
	StatusCode int
	Users      *[]models.Users
	StatusDesc string
}

type UsersAllCustomersDTO struct {
	StatusCode int
	Users      *[]interface{}
	StatusDesc string
}

type UsersBranchResponseDTO struct {
	StatusCode int
	Users      *[]models.Users
	StatusDesc string
}
