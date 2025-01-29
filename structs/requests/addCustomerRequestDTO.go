package requests

type AddCustomerRequestDTO struct {
	Name        string
	Category    string
	PhoneNumber string
	// ShopName            string
	IdType   string
	IdNumber string
	Nickname string
	// ShopAssistantName   string
	// ShopAssistantNumber string
	// Password string
	Email   string
	Dob     string
	AddedBy string
}

type UpdateCustomerRequestDTO struct {
	Name        string
	Category    string
	PhoneNumber string
	IdType      string
	IdNumber    string
	// ShopName            string
	Nickname string
	// ShopAssistantName   string
	// ShopAssistantNumber string
	// Password string
	Email   string
	Dob     string
	AddedBy string
}
