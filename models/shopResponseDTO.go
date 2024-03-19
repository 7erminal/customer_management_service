package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type ShopResponseDTO struct {
	StatusCode int    `orm: "omitempty"`
	Shop       *Shops `orm: "omitempty"`
	StatusDesc string `orm:"size(255)"`
}
