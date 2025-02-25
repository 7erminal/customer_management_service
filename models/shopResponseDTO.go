package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type ShopResponseDTO struct {
	StatusCode int
	Shop       *Shops
	StatusDesc string
}

type ShopsResponseDTO struct {
	StatusCode int
	Shop       []*interface{}
	StatusDesc string
}
