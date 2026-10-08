package responses

type IDTypeResponse struct {
	IdentificationTypeId string
	Name                 string
	Code                 string
}

type IDTypesResponseDTO struct {
	StatusCode int
	IdTypes    *[]IDTypeResponse
	StatusDesc string
}
