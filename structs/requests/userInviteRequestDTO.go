package requests

type UpdateUserInviteRequest struct {
	Email  string
	Status string
}

type UpdateTokenRequestDTO struct {
	UserId string
	Status string
}
