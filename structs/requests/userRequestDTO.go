package requests

type UserCredentialsDTO struct {
	Password string `orm:"size(255)"`
	Username string `orm:"size(255)"`
}

type SignUpCredDTO struct {
	Password string `orm:"size(255)"`
	Username string `orm:"size(255)"`
	AddedBy  string `orm:"size(255)"`
}

type SignUpDTO struct {
	Username     string
	Name         string
	Password     string
	Email        string
	Gender       string
	Dob          string
	PhoneNumber  string
	Role         string
	Branch       *string
	RoleRequired bool
	AddedBy      string
}

type UpdateUserPasswordRequest struct {
	UserId      int64  `validate:"required"`
	OldPassword string `validate:"required"`
	NewPassword string `validate:"required"`
}
