package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type CustomerResponseDTO struct {
	StatusCode int        `orm: "omitempty"`
	Customer   *Customers `orm: "omitempty"`
	StatusDesc string     `orm:"size(255)"`
}
