package requests

type ShopRequest struct {
	Name            string
	PhoneNumber     string
	Email           string
	ImageUrl        string
	Description     string
	Location        string
	AssistantName   string
	AssistantNumber string
}

type UpdateShopRequest struct {
	Name            string
	PhoneNumber     string
	Email           string
	ImageUrl        string
	Description     string
	Location        string
	AssistantName   string
	AssistantNumber string
}

type ShopBranchRequest struct {
	ShopId   string
	BranchId string
}
