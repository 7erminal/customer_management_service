package requests

type StringRequestDTO struct {
	Value string
}

type RegisterInviteRequestDTO struct {
	InviteBy string
	Email    string
	Role     string
	Link     string
}
