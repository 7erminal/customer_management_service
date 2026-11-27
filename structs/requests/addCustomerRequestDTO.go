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
	Branch      int64
	Status      int
}

type UpdateCustomerLastTxnRequest struct {
	TransactionDate string
}

type AddCustomerEmergencyContactRequestDTO struct {
	Name        string
	PhoneNumber string
	CustomerId  int64
}

type AddCustomerGuarantorRequestDTO struct {
	Name        string
	PhoneNumber string
	CustomerId  int64
}

type EditCustomerEmergencyContactRequestDTO struct {
	Name        string
	PhoneNumber string
}

type EditCustomerGuarantorRequestDTO struct {
	Name        string
	PhoneNumber string
}
