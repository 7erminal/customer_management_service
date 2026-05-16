package requests

type UpdateUserPasswordRequest struct {
	UserId      int64  `validate:"required"`
	OldPassword string `validate:"required"`
	NewPassword string `validate:"required"`
}
