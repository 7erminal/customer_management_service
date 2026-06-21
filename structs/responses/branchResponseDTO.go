package responses

import "time"

type BranchResp struct {
	BranchId     int64
	BranchName   string
	Description  string
	Location     string
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
