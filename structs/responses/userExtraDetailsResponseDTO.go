package responses

import "time"

type UserExtraDetails struct {
	UserDetailsId string
	Branch        *BranchResp
	Shop          *ShopResp
	Nickname      string
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     string
	ModifiedBy    string
	Active        int
}

type UserExtraDetailsResponseDTO struct {
	StatusCode  int
	UserDetails *UserExtraDetails
	StatusDesc  string
}
