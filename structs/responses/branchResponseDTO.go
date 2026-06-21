package responses

import "time"

type CurrencyResp struct {
	CurrencyId int64
	Symbol     string
	Currency   string
}

type CountryResp struct {
	CountryId   int64
	Country     string
	CountryCode string
	Currency    *CurrencyResp
}

type BranchResp struct {
	BranchId     int64
	BranchName   string
	Description  string
	Location     string
	Country      *CountryResp
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
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
