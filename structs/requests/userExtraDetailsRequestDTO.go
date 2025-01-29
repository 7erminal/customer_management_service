package requests

type AddUserExtraDetailsRequestDTO struct {
	Name                string
	Category            string
	PhoneNumber         string
	ShopName            string
	Nickname            string
	ShopAssistantName   string
	ShopAssistantNumber string
	// Password string
	Email   string
	Gender  string
	Dob     string
	AddedBy string
}
