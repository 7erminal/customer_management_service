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
	Branch      string
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
	Branch      string
	Status      string
}

type UpdateCustomerLastTxnRequest struct {
	TransactionDate string
}

type AddCustomerEmergencyContactRequestDTO struct {
	Name        string
	PhoneNumber string
	CustomerId  string
	AddedBy     string
}

type AddCustomerGuarantorRequestDTO struct {
	Name        string
	PhoneNumber string
	CustomerId  string
	AddedBy     string
}

type EditCustomerEmergencyContactRequestDTO struct {
	Name        string
	PhoneNumber string
	ModifiedBy  string
}

type EditCustomerGuarantorRequestDTO struct {
	Name        string
	PhoneNumber string
	ModifiedBy  string
}
