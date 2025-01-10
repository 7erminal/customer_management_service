package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type SignUpDTO struct {
	Name         string `orm:"size(255)"`
	Password     string `orm:"size(255)"`
	Email        string `orm:"size(255)"`
	Gender       string `orm:"size(255)"`
	Dob          string `orm:"size(255)"`
	PhoneNumber  string `orm:"size(255)"`
	Role         string `orm:"size(255)"`
	RoleRequired bool
}

func init() {
	orm.RegisterModel(new(SignUpDTO))
}
