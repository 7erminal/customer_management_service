package responses

type TokenDestructureResponseDTO struct {
	Email  string
	RoleId string
}

type InviteDecodeResponseDTO struct {
	StatusCode int
	Value      *TokenDestructureResponseDTO
	StatusDesc string
}

type StringResponseDTO struct {
	StatusCode int
	Value      string
	StatusDesc string
}
