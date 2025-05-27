package responses

type TokenDestructureResponseDTO struct {
	TokenId  string
	Email    string
	RoleId   string
	InviteBy string
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
