package responses

import (
	"customer_management_service/models"
	"time"
)

type InviteHashDTO struct {
	Token *models.UserTokens
}

type InviteHashResponseDTO struct {
	StatusCode int
	Value      *InviteHashDTO
	StatusDesc string
}

type UserInvitesResp struct {
	UserInviteId    int64
	InvitedBy       *models.Users
	InvitationToken *models.UserTokens
	Status          string
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
	Active          int
}

type UserInvitesResponseDTO struct {
	StatusCode  int
	UserInvites *[]interface{}
	StatusDesc  string
}

type UserInviteResponseDTO struct {
	StatusCode int
	UserInvite *models.UserInvites
	StatusDesc string
}
