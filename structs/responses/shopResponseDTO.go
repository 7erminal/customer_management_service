package responses

import "time"

type ShopBranchResp struct {
	ShopBranch BranchResp
	ShopId     string
	BranchId   string
}

type ShopResp struct {
	ShopId              string
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	PhoneNumber         string
	Email               string
	Image               string
	ShopLocation        string
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
	Active              int
	ShopBranches        []ShopBranchResp
}

type ShopListResponseDTO struct {
	Shops []ShopResp `json:"shops"`
}

type ShopResponse struct {
	StatusCode int
	StatusDesc string
	Result     ShopResp
}

type ShopsResponse struct {
	StatusCode int
	StatusDesc string
	Result     []ShopResp
}

type ShopBranchApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     struct {
		Branch BranchResp
		Shop   ShopResp
	}
}
