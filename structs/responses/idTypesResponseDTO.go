package responses

type IDTypeResponse struct {
	IdentificationTypeId int64
	Name                 string
	Code                 string
}

type IDTypesResponseDTO struct {
	StatusCode int
	IdTypes    *[]interface{}
	StatusDesc string
}
