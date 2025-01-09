package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type UserResponseDTO struct {
	StatusCode int
	User       *Users
	StatusDesc string
}
