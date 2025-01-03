package responses

type BranchesResponseDTO struct {
	StatusCode int
	Branches   *[]interface{}
	StatusDesc string
}
