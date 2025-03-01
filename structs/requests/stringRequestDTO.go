package requests

type StringRequestDTO struct {
	Value string
}

type RegisterInviteRequestDTO struct {
	InviteBy string
	Email    string
	Role     string
	Message  string
	Subject  string
	Links    []*string
}
