package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type CustomerResponseDTO struct {
	StatusCode int
	Customer   *Customers
	StatusDesc string
}
