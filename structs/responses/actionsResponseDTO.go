package responses

import "time"

type Actions struct {
	ActionId     string
	Action       string
	Description  string
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
	Active       int
}

type ActionResponseDTO struct {
	StatusCode int
	Action     *Actions
	StatusDesc string
}

type ActionsResponseDTO struct {
	StatusCode int
	Actions    *[]Actions
	StatusDesc string
}

type ActionsAllResponseDTO struct {
	StatusCode int
	Actions    *[]interface{}
	StatusDesc string
}
