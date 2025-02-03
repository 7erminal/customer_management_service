package requests

type AddCustomerRequestDTO struct {
	Name        string
	Category    string
	PhoneNumber string
	IdType      string
	IdNumber    string
	Nickname    string
	Location    string
	Email       string
	Dob         string
	AddedBy     string
}

type UpdateCustomerRequestDTO struct {
	Name        string
	Category    string
	PhoneNumber string
	IdType      string
	IdNumber    string
	Nickname    string
	Location    string
	Email       string
	Dob         string
	AddedBy     string
}
