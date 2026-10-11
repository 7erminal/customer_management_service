package responses

import "time"

type CurrencyResp struct {
	CurrencyId string
	Symbol     string
	Currency   string
}

type CountryResp struct {
	CountryId   string
	Country     string
	CountryCode string
	Currency    *CurrencyResp
}

type BranchResp struct {
	BranchId     string
	BranchName   string
	Description  string
	PhoneNumber  string
	Location     string
	Country      *CountryResp
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

type BranchResponseDTO struct {
	StatusCode int
	Result     *BranchResp
	StatusDesc string
}

type BranchesResponseDTO struct {
	StatusCode int
	Result     *[]BranchResp
	StatusDesc string
}
